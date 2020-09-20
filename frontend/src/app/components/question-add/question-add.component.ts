import { Component, OnInit } from '@angular/core';
import {ActivatedRoute, Router} from '@angular/router';
import {QuestionsService} from '../../services/questions.service';
import {HttpResponse} from '@angular/common/http';
import {QuestionsResponse} from '../../models/questions';

@Component({
  selector: 'app-question-add',
  templateUrl: './question-add.component.html',
  styleUrls: ['./question-add.component.css']
})
export class QuestionAddComponent implements OnInit {

  surveyId: string;
  constructor(
    public router: Router,
    private questionsService: QuestionsService,
    private activatedRoute: ActivatedRoute,
    ) { }

  ngOnInit(): void {
    this.activatedRoute.paramMap.subscribe(params => {
      this.surveyId = params.get('surveyId');
    });
  }
  onQuestionSubmit(surveyData): void{
    this.questionsService.postQuestion(
      surveyData.title,
      surveyData.text,
      surveyData.first,
      this.surveyId,
    ).subscribe((response: HttpResponse<QuestionsResponse>) => {
      console.log(response);
      if (response.status === 200) {
        this.router.navigate(['/surveys/edit', this.surveyId]);
      }
    });
  }
}
