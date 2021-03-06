import {Component, EventEmitter, Input, OnInit, Output} from '@angular/core';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';
import {ActivatedRoute, Router} from '@angular/router';
import {SurveysService} from '../../services/surveys.service';
import {Survey, SurveyResponse} from '../../models/survey';
import {HttpResponse} from '@angular/common/http';

@Component({
  selector: 'app-survey-form',
  templateUrl: './survey-form.component.html',
  styleUrls: ['./survey-form.component.css']
})
export class SurveyFormComponent implements OnInit {

  @Input() getSurvey: boolean;
  @Output() formData = new EventEmitter<any>();
  surveyForm: FormGroup;
  surveyId: string;
  survey: Survey;
  constructor(
    public router: Router,
    private surveysService: SurveysService,
    private formBuilder: FormBuilder,
    private activatedRoute: ActivatedRoute,
  ) {
    this.surveyForm = this.formBuilder.group({
    title: ['', [Validators.required]],
    summary: ['', [Validators.required]],
    disclaimer: ['', [Validators.required]],
    });
  }

  ngOnInit(): void {
    this.surveyForm.reset();
    if (this.getSurvey === true) {
      this.activatedRoute.paramMap.subscribe(params => {
        this.surveyId = params.get('surveyId');
        this.surveysService.getSurvey(this.surveyId).subscribe((response: HttpResponse<SurveyResponse>) => {
          if (response.status === 200) {
            this.survey = response.body.survey;
            console.log(this.survey);
            this.surveyForm.setValue({
              title: this.survey.title,
              summary: this.survey.summary,
              disclaimer: this.survey.disclaimer,
            });
          }
        });
      });
    }
  }
  onSurveySubmit(surveyData): void{
    this.formData.emit(surveyData);
  }
}
