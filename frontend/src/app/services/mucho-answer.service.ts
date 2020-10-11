import { Injectable } from '@angular/core';
import {HttpClient, HttpResponse} from '@angular/common/http';
import {Router} from '@angular/router';
import {Observable} from 'rxjs';
import {MultipleChoiceAnswer} from '../models/mucho';
import {Question} from '../models/questions';

@Injectable({
  providedIn: 'root'
})
export class MuchoAnswerService {

  constructor(
    private http: HttpClient,
    private router: Router,
  ) { }
  postAnswer(
    email: string,
    text: string,
    question: Question): Observable<HttpResponse<any>> {
    return this.http.post(`/api/answer/multiplechoice`, {
      email,
      text,
      question,
    }, {observe: 'response'});
  }
}
