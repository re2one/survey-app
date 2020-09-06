import { Component, OnInit } from '@angular/core';
import { SurveysService} from '../../services/surveys.service';
import {LoginService} from '../../services/login.service';
import {Survey, Surveys} from '../../models/survey';
import {Observable, of} from 'rxjs';
import {catchError, map} from 'rxjs/operators';
import {Router} from '@angular/router';

@Component({
  selector: 'app-surveys',
  templateUrl: './surveys.component.html',
  styleUrls: ['./surveys.component.css']
})
export class SurveysComponent implements OnInit {
  localSurveys: Array<Survey>;

  constructor(
    private surveysService: SurveysService,
    private loginService: LoginService,
    public router: Router,
  ) {
    this.localSurveys = [];
  }

  ngOnInit(): void {
    this.surveysService.getSurveys().subscribe( obj => {
      obj.surveys.forEach(survey => {
        this.localSurveys.push(survey);
      });
    });
    console.log(this.localSurveys);
  }

  permissionCheck(): boolean {
    const role = localStorage.getItem('role');
    return role === 'admin';
  }
  moveToAddForm(): void {
    this.router.navigate(['/surveys/survey-add']);
  }

}
