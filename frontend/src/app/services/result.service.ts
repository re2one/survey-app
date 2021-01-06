import {Injectable} from '@angular/core';
import {HttpClient, HttpResponse} from '@angular/common/http';
import {Router} from '@angular/router';
import {Observable} from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ResultService {

  constructor(
    private http: HttpClient,
    private router: Router,
  ) {
  }

  getResult(surveyId: number): Observable<HttpResponse<any>> {
    return this.http.get(`/api/results/${surveyId}`, {observe: 'response'});
  }

  getSingleResult(surveyId: string, email: string, questionId: string): Observable<HttpResponse<any>> {
    return this.http.get(`/api/results/${surveyId}/${email}/${questionId}`, {observe: 'response'});
  }
}
