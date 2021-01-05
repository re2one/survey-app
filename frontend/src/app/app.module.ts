import {BrowserModule} from '@angular/platform-browser';
import {NgModule} from '@angular/core';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';

import {AppRoutingModule} from './app-routing.module';
import {AppComponent} from './app.component';
import {LoginComponent} from './components/login/login.component';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';

import {MatTabsModule} from '@angular/material/tabs';
import {MatToolbarModule} from '@angular/material/toolbar';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatSelectModule} from '@angular/material/select';
import {MatInputModule} from '@angular/material/input';
import {MatButtonModule} from '@angular/material/button';
import {MatMenuModule} from '@angular/material/menu';
import {MatCardModule} from '@angular/material/card';
import {DeleteDialogComponent, SurveysComponent} from './components/surveys/surveys.component';
import {AuthInterceptor} from './services/auth-interceptor.service';
import {HTTP_INTERCEPTORS, HttpClientModule} from '@angular/common/http';

import {environment} from '../environments/environment';
import {SurveyFormComponent} from './components/survey-form/survey-form.component';
import {SurveyAddComponent} from './components/survey-add/survey-add.component';
import {SurveyEditComponent} from './components/survey-edit/survey-edit.component';
import {QuestionsComponent} from './components/questions/questions.component';
import {SurveyDetailsComponent} from './components/survey-details/survey-details.component';
import {QuestionAddComponent} from './components/question-add/question-add.component';
import {QuestionEditComponent} from './components/question-edit/question-edit.component';
import {QuestionDetailsComponent} from './components/question-details/question-details.component';
import {QuestionFormComponent} from './components/question-form/question-form.component';
import {MultipleFormComponent} from './components/multiple-form/multiple-form.component';
import {MultipleEditComponent} from './components/multiple-edit/multiple-edit.component';
import {MultipleAddComponent} from './components/multiple-add/multiple-add.component';
import {SurveyMainComponent, TimerAlertDialogComponent} from './components/survey-main/survey-main.component';
import {QuestionAnswerComponent} from './components/question-answer/question-answer.component';
import {MultipleAnswerComponent} from './components/multiple-answer/multiple-answer.component';
import {PuzzleAnswerComponent} from './components/puzzle-answer/puzzle-answer.component';
import {QuestionEditPuzzleComponent} from './components/question-edit-puzzle/question-edit-puzzle.component';
import {MatDialogModule} from '@angular/material/dialog';
import {PuzzleAddDialogComponent} from './components/puzzle-add-dialog/puzzle-add-dialog.component';
import {PuzzlePreviewComponent} from './components/puzzle-preview/puzzle-preview.component';
import {SurveyInspectComponent} from './components/survey-inspect/survey-inspect.component';


@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    SurveysComponent,
    SurveyFormComponent,
    SurveyAddComponent,
    SurveyEditComponent,
    QuestionsComponent,
    SurveyDetailsComponent,
    QuestionAddComponent,
    QuestionEditComponent,
    QuestionDetailsComponent,
    QuestionFormComponent,
    MultipleFormComponent,
    MultipleEditComponent,
    MultipleAddComponent,
    SurveyMainComponent,
    QuestionAnswerComponent,
    MultipleAnswerComponent,
    PuzzleAnswerComponent,
    QuestionEditPuzzleComponent,
    PuzzleAddDialogComponent,
    PuzzlePreviewComponent,
    TimerAlertDialogComponent,
    DeleteDialogComponent,
    SurveyInspectComponent,
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
    HttpClientModule,
    MatDialogModule,
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
