import {ChangeDetectorRef, Component, OnInit} from '@angular/core';
import {QuestionsService} from '../../services/questions.service';
import {Question, QuestionsResponse} from '../../models/questions';
import {HttpResponse} from '@angular/common/http';
import {ActivatedRoute, Router} from '@angular/router';
import {FullQuestionsService} from '../../services/full-questions.service';
import {MuchoAnswerService} from '../../services/mucho-answer.service';

@Component({
  selector: 'app-question-answer',
  templateUrl: './question-answer.component.html',
  styleUrls: ['./question-answer.component.css']
})
export class QuestionAnswerComponent implements OnInit {

  public order: number;
  public question: Question;
  public questionId: string;
  public surveyId: string;

  constructor(
    private cdr: ChangeDetectorRef,
    private questionsService: QuestionsService,
    private fullQuestuionsService: FullQuestionsService,
    private muchoAnswerService: MuchoAnswerService,
    private activatedRoute: ActivatedRoute,
    private router: Router
  ) {
  }

  ngOnInit(): void {
    this.activatedRoute.paramMap.subscribe(params => {
      this.questionId = params.get('questionId');
      this.surveyId = params.get('surveyId');
      this.order = parseInt(params.get('order'), 10);
      this.questionsService.getQuestion(this.questionId).subscribe((response: HttpResponse<QuestionsResponse>) => {
        if (response.status === 200) {
          this.question = response.body.question;
          console.log(this.question.type);
        }
      });
    });
  }
  onAnswerSubmit(answer): void {
    const email = localStorage.getItem('email');
    this.muchoAnswerService.postAnswer(email, answer.answer, this.question).subscribe((response: HttpResponse<QuestionsResponse>) => {
      if (response.status === 200) {
        this.fullQuestuionsService.postFullQuestion(
          email,
          this.question,
          this.order,
        ).subscribe((response2: HttpResponse<QuestionsResponse>) => {
          if (response2.status === 200) {
            console.log(response2.body);
            this.router.navigate(['survey', this.surveyId]);
          }
        });
      }
    });
  }
}
