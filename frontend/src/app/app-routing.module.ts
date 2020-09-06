import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { LoginComponent } from './components/login/login.component';
import { SurveysComponent } from './components/surveys/surveys.component';
import {SurveyFormComponent} from './components/survey-form/survey-form.component';
import {AuthGuardService} from './services/auth-guard.service';
import {RoleGuardService} from './services/role-guard.service';


const routes: Routes = [
  {path: '', redirectTo: 'login', pathMatch: 'full'},
  {path: 'login', component: LoginComponent},
  {path: 'surveys', canActivate: [AuthGuardService], component: SurveysComponent},
  {path: 'survey-form', canActivate: [AuthGuardService, RoleGuardService], component: SurveyFormComponent},
  {path: '**', redirectTo: 'login'}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
