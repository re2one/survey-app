import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { LoginComponent } from './components/login/login.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import {MatTabsModule} from '@angular/material/tabs';
import {MatToolbarModule} from '@angular/material/toolbar';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatSelectModule} from '@angular/material/select';
import {MatInputModule} from '@angular/material/input';
import {MatButtonModule} from '@angular/material/button';
import {MatMenu, MatMenuModule} from '@angular/material/menu';
import {MatCardModule} from '@angular/material/card';
import { SurveysComponent } from './components/surveys/surveys.component';
import {AuthInterceptor} from './services/auth-interceptor.service';
import {HTTP_INTERCEPTORS, HttpClientModule} from '@angular/common/http';

import { environment } from '../environments/environment';
import { SurveyFormComponent } from './components/survey-form/survey-form.component';
import { SurveyAddComponent } from './components/survey-add/survey-add.component';
import { SurveyEditComponent } from './components/survey-edit/survey-edit.component';
import { QuestionsComponent } from './components/questions/questions.component';
import { SurveyDetailsComponent } from './components/survey-details/survey-details.component';



@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    SurveysComponent,
    SurveyFormComponent,
    SurveyAddComponent,
    SurveyEditComponent,
    QuestionsComponent,
    SurveyDetailsComponent
  ],
  imports: [
    FormsModule,
    ReactiveFormsModule,
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    MatTabsModule,
    MatToolbarModule,
    MatFormFieldModule,
    MatSelectModule,
    MatInputModule,
    MatButtonModule,
    MatMenuModule,
    MatCardModule,
    HttpClientModule
  ],
  providers: [
    AppModule,
    {
      provide: 'BACKEND_API_URL',
      useValue: environment.backendUrl
    },
    {
      provide: HTTP_INTERCEPTORS,
      useClass: AuthInterceptor,
      multi: true
    },
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
