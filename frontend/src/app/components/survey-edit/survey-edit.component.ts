import { Component, OnInit } from '@angular/core';
import {HttpResponse} from '@angular/common/http';
import {Survey, SurveyResponse} from '../../models/survey';
import {Router} from '@angular/router';
import {SurveysService} from '../../services/surveys.service';
import {ActivatedRoute} from '@angular/router';

@Component({
  selector: 'app-survey-edit',
  templateUrl: './survey-edit.component.html',
  styleUrls: ['./survey-edit.component.css']
})
export class SurveyEditComponent implements OnInit {
  surveyId: string;
  survey: Survey;
  constructor(
    public router: Router,
    private surveysService: SurveysService,
    private activatedRoute: ActivatedRoute) { }

  ngOnInit(): void {
    this.activatedRoute.paramMap.subscribe(params => {
      this.surveyId = params.get('surveyId');
    });
  }
  onSurveySubmit(surveyData): void{
    this.surveysService.putSurvey(
      parseInt(this.surveyId, 10),
      surveyData.title,
      surveyData.summary,
      surveyData.introduction,
      surveyData.disclaimer
    ).subscribe((response: HttpResponse<SurveyResponse>) => {
      if (response.status === 200) {
        this.router.navigate(['/surveys']);
      }
    });
  }
}
