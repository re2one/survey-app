import {Injectable} from '@angular/core';
import {Observable} from 'rxjs';
import {HttpClient, HttpResponse} from '@angular/common/http';
import {Router} from '@angular/router';
import {Puzzlepiece} from '../models/puzzle';

@Injectable({
  providedIn: 'root'
})
export class QuestionsService {

  constructor(
    private http: HttpClient,
    private router: Router,
  ) {
  }

  getQuestions(surveyId: string): Observable<HttpResponse<any>> {
    return this.http.get(`/api/questions/${surveyId}`, {observe: 'response'});
  }

  getQuestion(questionId): Observable<any> {
    return this.http.get(`/api/questions/single/${questionId}`, {observe: 'response'});
  }

  getAnsweredQuestions(email: string, surveyid: string): Observable<any> {

    console.log('SURVEY ID: ' + surveyid);
    return this.http.get(`/api/questions/answered/${email}/${surveyid}`, {observe: 'response'});
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
    surveyId: number,
    title: string,
    first: string,
    text: string,
    type: string,
    bracket: string,
    next: string,
    secondToNext: string,
    typeOfNextQuestion: string): Observable<HttpResponse<any>> {
    return this.http.put(`/api/questions`, {
      ID,
      surveyid: surveyId,
      title,
      text,
      first,
      Survey: null,
      type,
      bracket,
      next,
      secondToNext,
      typeOfNextQuestion,
    }, {observe: 'response'});
  }

  answerPuzzle(pieces: Array<Puzzlepiece>): Observable<HttpResponse<any>> {
    return this.http.post(`/api/answer/puzzle`, pieces, {observe: 'response'});
  }
}
