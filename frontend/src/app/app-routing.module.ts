import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { LoginComponent } from './components/login/login.component';
import { SurveysComponent } from './components/surveys/surveys.component';
import {SurveyFormComponent} from './components/survey-form/survey-form.component';
import {SurveyAddComponent} from './components/survey-add/survey-add.component';
import {AuthGuardService} from './services/auth-guard.service';
import {RoleGuardService} from './services/role-guard.service';
import {SurveyEditComponent} from './components/survey-edit/survey-edit.component';


const routes: Routes = [
  {path: '', redirectTo: 'login', pathMatch: 'full'},
  {path: 'login', component: LoginComponent},
  {path: 'surveys', canActivate: [AuthGuardService], component: SurveysComponent},
  {path: 'surveys/survey-add', canActivate: [AuthGuardService, RoleGuardService], component: SurveyAddComponent},
  {path: 'surveys/survey-edit/:surveyId', canActivate: [AuthGuardService, RoleGuardService], component: SurveyEditComponent},
  {path: '**', redirectTo: 'login'}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
