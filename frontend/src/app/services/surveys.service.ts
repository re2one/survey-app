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
  getSurvey(surveyId): Observable<any> {
    return this.http.get(`/api/surveys/${surveyId}`, {observe: 'response'});
  }
  postSurvey(
    title: string,
    summary: string,
    introduction: string,
    disclaimer: string): Observable<HttpResponse<any>> {
    return this.http.post(`/api/surveys`, {
      title,
      summary,
      introduction,
      disclaimer
    }, {observe: 'response'});
  }
  deleteSurvey(id: number): Observable<HttpResponse<any>> {
    return this.http.delete(`/api/surveys/${id}`, {observe: 'response'});
  }
  putSurvey(
    ID: number,
    title: string,
    summary: string,
    introduction: string,
    disclaimer: string): Observable<HttpResponse<any>> {
    return this.http.put(`/api/surveys`, {
      ID,
      title,
      summary,
      introduction,
      disclaimer
    }, {observe: 'response'});
  }
}
