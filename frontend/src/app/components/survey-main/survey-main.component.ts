import {ChangeDetectorRef, Component, EventEmitter, Inject, OnInit, Output} from '@angular/core';
import {ActivatedRoute, Router} from '@angular/router';
import {FullQuestionsService} from '../../services/full-questions.service';
import {HttpResponse} from '@angular/common/http';
import {FullQuestion, FullQuestions} from '../../models/questions';
import {MAT_DIALOG_DATA, MatDialog,} from '@angular/material/dialog';
import {ResultService} from '../../services/result.service';

@Component({
  selector: 'app-survey-main',
  templateUrl: './survey-main.component.html',
  styleUrls: ['./survey-main.component.css']
})
export class SurveyMainComponent implements OnInit {
  private surveyId: string;
  public questions: Array<FullQuestion>;
  public finished: boolean;
  public average: number;

  constructor(
    public dialog: MatDialog,
    private router: Router,
    private fullQuestionService: FullQuestionsService,
    private activatedRoute: ActivatedRoute,
    private resultService: ResultService,
    private cdr: ChangeDetectorRef,
  ) {
    this.average = 0.0;
  }

  ngOnInit(): void {
    this.activatedRoute.paramMap.subscribe(params => {
      this.surveyId = params.get('surveyId');
      const email = localStorage.getItem('email');
      this.fullQuestionService.getFullQuestions(this.surveyId, email).subscribe((response: HttpResponse<FullQuestions>) => {
        this.questions = response.body.questions;
        this.finished = response.body.finished;
        if (this.finished === true) {
          this.resultService.getAverage(this.surveyId, localStorage.getItem('email')).subscribe(response2 => {
            this.average = response2.body;
            this.cdr.detectChanges();
          });
        }
      });
    });
  }

  moveToAnswer(questionId: number): void {
    this.router.navigate(['question/answer', this.surveyId, questionId.toString(10), this.questions.length]);
  }

  questionHelper(question: FullQuestion): void {
    switch (question.type) {
      case 'puzzle': {
        const example = question.example === 'true';
        this.openTimerAlert(question.questionId, question.type, example);
        break;
      }
      default: {
        const example = question.example === 'true';
        if (example) {
          this.openTimerAlert(question.questionId, question.type, example);
          break;
        }
        this.moveToAnswer(question.questionId);
        break;
      }
    }
  }

  openTimerAlert(questionid: number, typeOfQuestion: string, example: boolean): void {
    const dialogRef = this.dialog.open(TimerAlertDialogComponent, {
      data: {
        typeOfQuestion,
        example
      }
    });
    dialogRef.componentInstance.shouldProceed.subscribe(event => {
      if (event) {
        this.moveToAnswer(questionid);
      }
    });
  }
}

@Component({
  selector: 'app-timer-alert-dialog',
  template: `
    <h1 mat-dialog-title>Attention</h1>
    <div mat-dialog-content>
      <p *ngIf="data.typeOfQuestion==='puzzle'">
        After confirmation of this dialog, a timer of 15 seconds will start.
        <br>
        Please try to memorize the game state during this time and replicate it afterwards.
      </p>
      <p *ngIf="data.typeOfQuestion==='multiplechoice'">
        This is a single choice question.
        <br>
        Please choose the option that suits you best.
      </p>
      <p *ngIf="data.example && data.typeOfQuestion==='puzzle'">
        THIS IS AN EXAMPLE QUESTION!
        <br>
        Your score on this question will not be counted towards your total score.
      </p>
      <p *ngIf="data.example && data.typeOfQuestion==='multiplechoice'">
        THIS IS AN EXAMPLE QUESTION!
        <br>
        Your answer on this question will not impact your final result.
      </p>
    </div>
    <div mat-dialog-actions>
      <button mat-button mat-dialog-close (click)="emitProceeding(true)">Start</button>
      <button mat-button mat-dialog-close (click)="emitProceeding(false)">Cancel</button>
    </div>
  `,
})
export class TimerAlertDialogComponent implements OnInit {

  @Output() shouldProceed = new EventEmitter<boolean>();

  constructor(
    @Inject(MAT_DIALOG_DATA) public data: any
  ) {

  }

  ngOnInit(): void {
    console.log(this.data);
  }

  emitProceeding(action: boolean): void {
    this.shouldProceed.emit(action);
  }
}
