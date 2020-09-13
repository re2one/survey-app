import {Component, EventEmitter, Input, OnInit, Output} from '@angular/core';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';
import {Question, QuestionsResponse} from '../../models/questions';
import {AnswerResponse, Mucho} from '../../models/mucho';
import {ActivatedRoute, Router} from '@angular/router';
import {QuestionsService} from '../../services/questions.service';
import {MuchoService} from '../../services/mucho.service';
import {HttpResponse} from '@angular/common/http';

@Component({
  selector: 'app-multiple-form',
  templateUrl: './multiple-form.component.html',
  styleUrls: ['./multiple-form.component.css']
})
export class MultipleFormComponent implements OnInit {
  @Input() getAnswer: boolean;
  @Output() formData = new EventEmitter<any>();
  multipleForm: FormGroup;
  multipleId: string;
  answer: Mucho;
  constructor(
    public router: Router,
    private multipleService: MuchoService,
    private formBuilder: FormBuilder,
    private activatedRoute: ActivatedRoute
  ) {
    this.multipleForm = this.formBuilder.group({
      text: ['', [Validators.required]],
      questionId: [''],
    });
  }

  ngOnInit(): void {
    this.multipleForm.reset();
    if (this.getAnswer === true) {
      this.activatedRoute.paramMap.subscribe(params => {
        this.multipleId = params.get('answerId');
        this.multipleService.getAnswer(this.multipleId).subscribe((response: HttpResponse<AnswerResponse>) => {
          if (response.status === 200) {
            this.answer = response.body.choice;
            this.multipleForm.setValue({
              text: this.answer.text,
              questionId: this.answer.questionid
            });
          }
        });
      });
    }
  }
  onMultipleSubmit(multipleData): void{
    this.formData.emit(multipleData);
  }

}
