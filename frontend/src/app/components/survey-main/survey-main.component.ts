import {ChangeDetectorRef, Component, EventEmitter, OnInit, Output} from '@angular/core';
import {ActivatedRoute, Router} from '@angular/router';
import {FullQuestionsService} from '../../services/full-questions.service';
import {HttpResponse} from '@angular/common/http';
import {FullQuestion, FullQuestions} from '../../models/questions';
import {MatDialog} from '@angular/material/dialog';
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
        this.openTimerAlert(question.questionId);
        break;
      }
      default: {
        this.moveToAnswer(question.questionId);
        break;
      }
    }
  }

  openTimerAlert(questionid: number): void {
    const dialogRef = this.dialog.open(TimerAlertDialogComponent);
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
      After confirmation of this dialog, a timer of 15 seconds will start.
      <br>
      Please try to memorize the game state during this time and replicate it afterwards.
    </div>
    <div mat-dialog-actions>
      <button mat-button mat-dialog-close (click)="emitProceeding(true)">Start</button>
      <button mat-button mat-dialog-close (click)="emitProceeding(false)">Cancel</button>
    </div>
  `,
})
export class TimerAlertDialogComponent {
  @Output() shouldProceed = new EventEmitter<boolean>();

  emitProceeding(action: boolean): void {
    this.shouldProceed.emit(action);
  }
}
