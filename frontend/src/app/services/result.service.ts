import { Injectable } from '@angular/core';
import {HttpClient, HttpResponse} from '@angular/common/http';
import {Router} from '@angular/router';
import {Question} from '../models/questions';
import {Observable} from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ResultService {

  constructor(
    private http: HttpClient,
    private router: Router,
    ) { }

  getResult(surveyId: number): Observable<HttpResponse<any>> {
    return this.http.get(`/api/results/${surveyId}`, {observe: 'response'});
  }
}
