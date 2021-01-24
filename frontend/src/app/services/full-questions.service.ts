import {Injectable} from '@angular/core';
import {HttpClient, HttpResponse} from '@angular/common/http';
import {Observable} from 'rxjs';
import {Question} from '../models/questions';

@Injectable({
  providedIn: 'root'
})
export class FullQuestionsService {

  constructor(
    private http: HttpClient
  ) {
  }

  getFullQuestions(surveyId: string, email: string): Observable<HttpResponse<any>> {
    return this.http.get(`/api/fullquestions/${surveyId}/${email}`, {observe: 'response'});
  }

  postFullQuestion(email: string, question: Question, order: number): Observable<HttpResponse<any>> {
    return this.http.post(`/api/fullquestions/answered/${email}/${order}`, {
      ID: question.ID,
      surveyId: question.surveyid,
      title: question.title,
      text: question.text,
      first: question.first,
      Survey: null,
      type: question.type,
    }, {observe: 'response'});
  }

  postViewed(email: string, question: Question): Observable<HttpResponse<any>> {
    return this.http.post(`/api/fullquestions/viewed/${email}`, {
      ID: question.ID,
      surveyId: question.surveyid,
      title: question.title,
      text: question.text,
      first: question.first,
      Survey: null,
      type: question.type,
    }, {observe: 'response'});
  }
}
