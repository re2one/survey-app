import {Component, EventEmitter, Input, OnInit, Output} from '@angular/core';
import {Question} from '../../models/questions';
import {MuchoService} from '../../services/mucho.service';
import {HttpResponse} from '@angular/common/http';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';
import {SubmitDialogComponent} from '../submit-dialog/submit-dialog.component';
import {MatDialog} from '@angular/material/dialog';

@Component({
  selector: 'app-multiple-answer',
  templateUrl: './multiple-answer.component.html',
  styleUrls: ['./multiple-answer.component.css']
})
export class MultipleAnswerComponent implements OnInit {
  @Input() question: Question;
  @Output() answer = new EventEmitter<any>();
  public answerForm: FormGroup;
  public localAnswers: Map<any, any>;
  constructor(
    private muchoService: MuchoService,
    private formBuilder: FormBuilder,
    public dialog: MatDialog,
  ) {
    this.localAnswers = new Map();
    this.answerForm = this.formBuilder.group({
      answer: ['', [Validators.required]]
    });
  }

  ngOnInit(): void {
    if (this.question.type === 'multiplechoice'){
      this.muchoService.getAnswers(this.question.ID.toString(10)).subscribe((response: HttpResponse<any>) => {
        if (response.status === 200) {
          response.body.choices.forEach(answer => {
            this.localAnswers.set(answer.ID, answer);
          });
          console.log(this.localAnswers);
        }
      });
    }
  }

  openSubmitAlert(answer): void {
    const dialogRef = this.dialog.open(SubmitDialogComponent);
    dialogRef.componentInstance.shouldProceed.subscribe(event => {
      if (event) {
        this.onAnswerSubmit(answer);
      }
    });
  }

  onAnswerSubmit(answer): void {
    this.answer.emit(answer);
  }

}
