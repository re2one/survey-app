import { Injectable } from '@angular/core';
import {Observable} from 'rxjs';
import {HttpClient, HttpResponse} from '@angular/common/http';
import {Router} from '@angular/router';
import {Mucho} from '../models/mucho';

@Injectable({
  providedIn: 'root'
})
export class MuchoService {

  constructor(
    private http: HttpClient,
    private router: Router,
  ) { }

  getAnswers(questionId: string): Observable<HttpResponse<any>>  {
    return this.http.get<Mucho>(`/api/choices/${questionId}`, {observe: 'response'});
  }
  getAnswer(answerId): Observable<any> {
    return this.http.get(`/api/choices/single/${answerId}`, {observe: 'response'});
  }
  postAnswer(
    text: string,
    questionId: string): Observable<HttpResponse<any>> {
    return this.http.post(`/api/choices/${questionId}`, {
      text,
    }, {observe: 'response'});
  }
  deleteAnswers(id: number): Observable<HttpResponse<any>> {
    return this.http.delete(`/api/choices/${id}`, {observe: 'response'});
  }
  putAnswer(
    ID: number,
    questionId: string,
    text: string,
    nextQuestion: string,
  ): Observable<HttpResponse<any>> {
    return this.http.put(`/api/choices`, {
      ID,
      questionId,
      text,
      nextQuestion,
      Question: null,
    }, {observe: 'response'});
  }
}
