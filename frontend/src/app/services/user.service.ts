import {Injectable} from '@angular/core';
import {HttpClient, HttpResponse} from '@angular/common/http';
import {Observable} from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  constructor(
    private http: HttpClient,
  ) {
  }

  getAll(): Observable<HttpResponse<any>> {
    return this.http.get(`/api/users`, {observe: 'response'});
  }
}
