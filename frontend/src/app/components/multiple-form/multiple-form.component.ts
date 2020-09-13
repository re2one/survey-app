import {ChangeDetectorRef, Component, EventEmitter, Input, OnInit, Output} from '@angular/core';
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
  @Input() surveyId: string;
  @Output() formData = new EventEmitter<any>();
  multipleForm: FormGroup;
  multipleId: string;
  answer: Mucho;
  questionz: Array<SelectOptions>;
  constructor(
    public router: Router,
    private multipleService: MuchoService,
    private questionsService: QuestionsService,
    private formBuilder: FormBuilder,
    private activatedRoute: ActivatedRoute,
    private cdr: ChangeDetectorRef
  ) {
    this.multipleForm = this.formBuilder.group({
      text: ['', [Validators.required]],
      questionId: [''],
      nextQuestion: [''],
    });
    this.questionz = new Array <SelectOptions> ();
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
              questionId: this.answer.questionid,
              nextQuestion: this.answer.nextQuestion,
            });
          }
        });
        setTimeout(() => {
          this.questionsService.getQuestions(this.surveyId).subscribe( (response: HttpResponse<any>) => {
            if (response.status === 200) {
              response.body.questions.forEach(question => {
                const option = new SelectOptions(question.ID, question.title);
                this.questionz.push(option);
              });
              this.cdr.detectChanges();
            }
          });
        }, 0);
      });
    }
  }
  onMultipleSubmit(multipleData): void{
    this.formData.emit(multipleData);
  }

}

class SelectOptions {
  value: string;
  viewValue: string;
  constructor(value: string, viewValue: string) {
    this.value = value;
    this.viewValue = viewValue;
  }
}
