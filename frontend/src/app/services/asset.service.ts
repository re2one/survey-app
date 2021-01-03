import {Injectable} from '@angular/core';
import {HttpClient, HttpResponse} from '@angular/common/http';
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

  postIntroduction(
    fileToUpload: File,
    surveyId: string,
  ): Observable<HttpResponse<any>> {
    const endpoint = `/api/assets/introduction/${surveyId}`;
    const formData: FormData = new FormData();
    formData.append('fileKey', fileToUpload, fileToUpload.name);
    return this.http.post(endpoint, formData, {observe: 'response'});
  }

  getIntroduction(
    surveyId: string,
  ): Observable<HttpResponse<any>> {
    return this.http.get(`/api/assets/introduction`, {observe: 'response'});
  }

  getFilenames(
    surveyId: string,
    questionId: string,
  ): Observable<HttpResponse<any>> {
    return this.http.get(`/api/assets/${surveyId}/${questionId}`, {observe: 'response'});
  }
}
