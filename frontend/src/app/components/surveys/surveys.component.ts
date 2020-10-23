import {ChangeDetectorRef, Component, OnInit} from '@angular/core';
import { SurveysService} from '../../services/surveys.service';
import {LoginService} from '../../services/login.service';
import {Survey} from '../../models/survey';
import {Router} from '@angular/router';
import {HttpResponse} from '@angular/common/http';
import {BehaviorSubject} from 'rxjs';

@Component({
  selector: 'app-surveys',
  templateUrl: './surveys.component.html',
  styleUrls: ['./surveys.component.css']
})
export class SurveysComponent implements OnInit{
  localSurveys: Map<any, any>;

  constructor(
    private surveysService: SurveysService,
    private loginService: LoginService,
    private cdr: ChangeDetectorRef,
    public router: Router,
  ) {
    this.localSurveys = new Map();
  }

  ngOnInit(): void {
    setTimeout(() => {
      this.surveysService.getSurveys().subscribe( obj => {
        obj.surveys.forEach(survey => {
          this.localSurveys.set(survey.ID, survey);
        });
      });
    }, 0);
  }

  permissionCheck(): boolean {
    const role = localStorage.getItem('role');
    return role === 'admin';
  }
  moveToAddForm(): void {
    this.router.navigate(['/surveys/add']);
  }
  delete(id: number): void {
    this.surveysService.deleteSurvey(id).subscribe((response: HttpResponse<any>) => {
      if (response.status === 200) {
        this.localSurveys.delete(id);
        this.cdr.detectChanges();
      }
    });
  }
  moveToEditForm(surveyId: number): void {
    this.router.navigate(['/surveys/edit', surveyId]);
  }
  moveToDetails(surveyId: number): void{
    this.router.navigate(['/surveys/details', surveyId]);
  }
  getResult(surveyId: number): void{
    console.log('Result');
  }
}
