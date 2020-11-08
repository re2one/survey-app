import { Injectable } from '@angular/core';
import {HttpClient, HttpResponse} from '@angular/common/http';
import {Question} from '../models/questions';
import {Observable} from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AssetService {

  constructor(
    private http: HttpClient,
  ) { }
  addDirectory(
    surveyId: string,
    questionId: string,
  ): Observable<HttpResponse<any>> {
    return this.http.post(`/api/assets/directory/${surveyId}/${questionId}`, {
    }, {observe: 'response'});
  }
  postFile(
    fileToUpload: File,
    surveyId: string,
    questionId: string,
    ): Observable<HttpResponse<any>> {
    const endpoint = `/api/assets/upload/${surveyId}/${questionId}`;
    const formData: FormData = new FormData();
    formData.append('fileKey', fileToUpload, fileToUpload.name);
    return this.http.post(endpoint, formData, {observe: 'response'});
  }
}
