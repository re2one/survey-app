import { Component, OnInit } from '@angular/core';
import {HttpResponse} from '@angular/common/http';
import {Survey, SurveyResponse} from '../../models/survey';
import {SurveysService} from '../../services/surveys.service';
import {ActivatedRoute, Router} from '@angular/router';

@Component({
  selector: 'app-survey-details',
  templateUrl: './survey-details.component.html',
  styleUrls: ['./survey-details.component.css']
})
export class SurveyDetailsComponent implements OnInit {
  survey: Survey;
  surveyId: string;
  constructor(
    private surveysService: SurveysService,
    private activatedRoute: ActivatedRoute,
    private router: Router
  ) {
    this.survey = new Survey();
  }

  ngOnInit(): void {
    this.activatedRoute.paramMap.subscribe(params => {
      this.surveyId = params.get('surveyId');
      this.surveysService.getSurvey(this.surveyId).subscribe((response: HttpResponse<SurveyResponse>) => {
        if (response.status === 200) {
          this.survey = response.body.survey;
        }
      });
    });
  }
  permissionCheck(): boolean {
    const role = localStorage.getItem('role');
    return role === 'admin';
  }
  moveToMain(): void{
    this.router.navigate(['survey', this.surveyId]);
  }
}
