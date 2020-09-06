import { Injectable } from '@angular/core';
import {Observable} from 'rxjs';
import {HttpClient} from '@angular/common/http';
import {Router} from '@angular/router';
import {Survey, Surveys} from '../models/survey';

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
}
