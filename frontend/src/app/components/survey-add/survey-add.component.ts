import { Component, OnInit, Output} from '@angular/core';
import {HttpResponse} from '@angular/common/http';
import {SurveyResponse} from '../../models/survey';
import {Router} from '@angular/router';
import {SurveysService} from '../../services/surveys.service';
import {EventEmitter} from '@angular/core';

@Component({
  selector: 'app-survey-add',
  templateUrl: './survey-add.component.html',
  styleUrls: ['./survey-add.component.css']
})
export class SurveyAddComponent implements OnInit {
  constructor(
    public router: Router,
    private surveysService: SurveysService,
    ) { }

  ngOnInit(): void {
  }
  onSurveySubmit(surveyData): void{
    this.surveysService.postSurvey(
      surveyData.title,
      surveyData.summary,
      surveyData.introduction,
      surveyData.disclaimer
    ).subscribe((response: HttpResponse<SurveyResponse>) => {
      console.log(response);
      if (response.status === 200) {
        this.router.navigate(['/surveys']);
      }
    });
  }
}
