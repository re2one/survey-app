import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { LoginComponent } from './components/login/login.component';
import { SurveysComponent } from './components/surveys/surveys.component';
import {SurveyFormComponent} from './components/survey-form/survey-form.component';
import {SurveyAddComponent} from './components/survey-add/survey-add.component';
import {AuthGuardService} from './services/auth-guard.service';
import {RoleGuardService} from './services/role-guard.service';
import {SurveyEditComponent} from './components/survey-edit/survey-edit.component';
import {QuestionsComponent} from './components/questions/questions.component';
import {SurveyDetailsComponent} from './components/survey-details/survey-details.component';


const routes: Routes = [
  {path: '', redirectTo: 'login', pathMatch: 'full'},
  {path: 'login', component: LoginComponent},
  {path: 'surveys', canActivate: [AuthGuardService], component: SurveysComponent},
  {path: 'surveys/add', canActivate: [AuthGuardService, RoleGuardService], component: SurveyAddComponent},
  {path: 'surveys/edit/:surveyId', canActivate: [AuthGuardService, RoleGuardService], component: SurveyEditComponent},
  {path: 'surveys/details/:surveyId', canActivate: [AuthGuardService], component: SurveyDetailsComponent},
  {path: '**', redirectTo: 'login'}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
