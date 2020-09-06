import { Injectable } from '@angular/core';
import {Observable} from 'rxjs';
import {HttpClient, HttpResponse} from '@angular/common/http';
import {Router} from '@angular/router';
import {Survey, SurveyResponse, Surveys} from '../models/survey';

@Injectable({
  providedIn: 'root'
})
export class SurveysService {

  constructor(
    private http: HttpClient,
    private router: Router,
  ) { }

  getSurveys(): Observable<Surveys> {
    return this.http.get<Surveys>(`/api/surveys`);
  }
  postSurvey(
    title: string,
    summary: string,
    description: string,
    disclaimer: string): Observable<HttpResponse<any>> {
    return this.http.post(`/api/surveys`, {
      title,
      summary,
      description,
      disclaimer
    }, {observe: 'response'});
  }
}
