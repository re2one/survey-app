import { Injectable } from '@angular/core';
import {Observable} from 'rxjs';
import {Surveys} from '../models/survey';
import {HttpClient, HttpResponse} from '@angular/common/http';
import {Router} from '@angular/router';

@Injectable({
  providedIn: 'root'
})
export class QuestionsService {

  constructor(
    private http: HttpClient,
    private router: Router,
  ) { }
  getQuestions(surveyId: string): Observable<HttpResponse<any>> {
    return this.http.get(`/api/questions/${surveyId}`, {observe: 'response'});
  }
}
