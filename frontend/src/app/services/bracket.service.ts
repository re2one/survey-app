import { Injectable } from '@angular/core';
import {HttpClient, HttpResponse} from '@angular/common/http';
import {Observable} from 'rxjs';
import {Mucho} from '../models/mucho';

@Injectable({
  providedIn: 'root'
})
export class BracketService {

  constructor(
    private http: HttpClient,
  ) { }
  getBrackets(surveyId: string): Observable<HttpResponse<any>>  {
    return this.http.get(`/api/brackets/${surveyId}`, {observe: 'response'});
  }
  postBracket(surveyId: string, bracketName: string): Observable<HttpResponse<any>>  {
    return this.http.post(`/api/brackets/${surveyId}`, {name: bracketName}, {observe: 'response'});
  }
}
