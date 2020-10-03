import {ChangeDetectorRef, Component, Input, OnInit} from '@angular/core';
import {ActivatedRoute, Router} from '@angular/router';
import {FullQuestionsService} from '../../services/full-questions.service';
import {HttpResponse} from '@angular/common/http';
import {SurveyResponse} from '../../models/survey';
import {FullQuestion, FullQuestions} from '../../models/questions';

@Component({
  selector: 'app-survey-main',
  templateUrl: './survey-main.component.html',
  styleUrls: ['./survey-main.component.css']
})
export class SurveyMainComponent implements OnInit {
  private surveyId: string;
  public questions: Array<FullQuestion>;
  constructor(
    private router: Router,
    private fullQuestionService: FullQuestionsService,
    private activatedRoute: ActivatedRoute,
    private cdr: ChangeDetectorRef,
  ) { }

  ngOnInit(): void {
    this.activatedRoute.paramMap.subscribe(params => {
      this.surveyId = params.get('surveyId');
      const email = localStorage.getItem('email');
      this.fullQuestionService.getFullQuestions(this.surveyId, email).subscribe((response: HttpResponse<FullQuestions>) => {
        this.questions = response.body.questions;
      });
    });
  }
  moveToAnswer(questionId: number): void {
    this.router.navigate(['question/answer', this.surveyId, questionId.toString(10)]);
  }
}
