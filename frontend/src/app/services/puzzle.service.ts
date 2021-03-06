import {Injectable} from '@angular/core';
import {HttpClient, HttpResponse} from '@angular/common/http';
import {Observable} from 'rxjs';
import {PuzzlePiece} from '../components/question-edit-puzzle/question-edit-puzzle.component';

@Injectable({
  providedIn: 'root'
})
export class PuzzleService {

  constructor(
    private http: HttpClient,
  ) { }

  update(
    surveyId: string,
    questionId: string,
    puzzlepieces: Array<PuzzlePiece>
  ): Observable<HttpResponse<any>> {
    return this.http.put(`/api/puzzle/${surveyId}/${questionId}`, puzzlepieces, {observe: 'response'});
  }

  getAll(questionId: string): Observable<HttpResponse<any>> {
    return this.http.get(`/api/puzzle/${questionId}`, {observe: 'response'});
  }

  getAllForQuestionaire(questionId: string, email: string): Observable<HttpResponse<any>> {
    return this.http.get(`/api/puzzle/${questionId}/${email}`, {observe: 'response'});
  }
}
