import { Injectable } from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Observable} from 'rxjs';


@Injectable({
  providedIn: 'root'
})
export class AuthTestService {

  constructor(
    private http: HttpClient
  ) { }
  get(): Observable<{ message: string }> {
    return this.http.get<{message: string}>(`/api/surveys`);
  }
}
