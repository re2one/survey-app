import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { LoginComponent } from './components/login/login.component';
import { SurveysComponent } from './components/surveys/surveys.component';
import {SurveyAddComponent} from './components/survey-add/survey-add.component';
import {AuthGuardService} from './services/auth-guard.service';
import {RoleGuardService} from './services/role-guard.service';
import {SurveyEditComponent} from './components/survey-edit/survey-edit.component';
import {SurveyDetailsComponent} from './components/survey-details/survey-details.component';
import {QuestionEditComponent} from './components/question-edit/question-edit.component';
import {QuestionAddComponent} from './components/question-add/question-add.component';


const routes: Routes = [
  {path: '', redirectTo: 'login', pathMatch: 'full'},
  {path: 'login', component: LoginComponent},
  {path: 'surveys', canActivate: [AuthGuardService], component: SurveysComponent},
  {path: 'surveys/add', canActivate: [AuthGuardService, RoleGuardService], component: SurveyAddComponent},
  {path: 'surveys/edit/:surveyId', canActivate: [AuthGuardService, RoleGuardService], component: SurveyEditComponent},
  {path: 'surveys/details/:surveyId', canActivate: [AuthGuardService], component: SurveyDetailsComponent},
  {path: 'questions/add/:surveyId', canActivate: [AuthGuardService, RoleGuardService], component: QuestionAddComponent},
  {path: 'questions/edit/:questionId', canActivate: [AuthGuardService, RoleGuardService], component: QuestionEditComponent},
  {path: '**', redirectTo: 'login'}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
