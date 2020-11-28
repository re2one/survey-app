import {Injectable} from '@angular/core';
import {Observable} from 'rxjs';
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
  getQuestion(questionId): Observable<any> {
    return this.http.get(`/api/questions/single/${questionId}`, {observe: 'response'});
  }
  postQuestion(
    title: string,
    text: string,
    first: string,
    surveyId: string,
    type: string,
    ): Observable<HttpResponse<any>> {
    return this.http.post(`/api/questions/${surveyId}`, {
      title,
      text,
      first,
      type,
    }, {observe: 'response'});
  }

  deleteQuestions(id: number): Observable<HttpResponse<any>> {
    return this.http.delete(`/api/questions/${id}`, {observe: 'response'});
  }

  putQuestion(
    ID: number,
    surveyId: string,
    title: string,
    first: string,
    text: string,
    type: string,
    bracket: string): Observable<HttpResponse<any>> {
    return this.http.put(`/api/questions`, {
      ID,
      surveyId,
      title,
      text,
      first,
      Survey: null,
      type,
      bracket
    }, {observe: 'response'});
  }
}
