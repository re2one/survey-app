import {Component, EventEmitter, Input, OnInit, Output} from '@angular/core';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';
import {Question, QuestionsResponse} from '../../models/questions';
import {ActivatedRoute, Router} from '@angular/router';
import {QuestionsService} from '../../services/questions.service';
import {HttpResponse} from '@angular/common/http';

@Component({
  selector: 'app-question-form',
  templateUrl: './question-form.component.html',
  styleUrls: ['./question-form.component.css']
})
export class QuestionFormComponent implements OnInit {

  @Input() getQuestion: boolean;
  @Input() brackets: Array<any>;
  @Output() formData = new EventEmitter<any>();
  questionForm: FormGroup;
  questionId: string;
  question: Question;
  constructor(
    public router: Router,
    private questionsService: QuestionsService,
    private formBuilder: FormBuilder,
    private activatedRoute: ActivatedRoute,
  ) {
    this.questionForm = this.formBuilder.group({
      title: ['', [Validators.required]],
      text: ['', [Validators.required]],
      first: ['', [Validators.required]],
      type: ['', [Validators.required]],
      bracket: ['', [Validators.required]],
      example: ['', [Validators.required]],
      surveyId: [''],
    });
  }

  ngOnInit(): void {
    this.questionForm.reset();
    if (this.getQuestion === true) {
      this.activatedRoute.paramMap.subscribe(params => {
        this.questionId = params.get('questionId');
        this.questionsService.getQuestion(this.questionId).subscribe((response: HttpResponse<QuestionsResponse>) => {
          if (response.status === 200) {
            this.question = response.body.question;
            this.questionForm.setValue({
              title: this.question.title,
              text: this.question.text,
              surveyId: this.question.surveyid,
              first: this.question.first,
              type: this.question.type,
              bracket: this.question.bracket || 'none',
              example: this.question.example,
            });
          }
        });
      });
    }
  }
  onQuestionSubmit(questionData): void{
    this.formData.emit(questionData);
  }
}
