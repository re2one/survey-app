import {Component, OnInit, Output} from '@angular/core';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';
import {Router} from '@angular/router';
import {SurveysService} from '../../services/surveys.service';
import {EventEmitter} from '@angular/core';

@Component({
  selector: 'app-survey-form',
  templateUrl: './survey-form.component.html',
  styleUrls: ['./survey-form.component.css']
})
export class SurveyFormComponent implements OnInit {

  @Output() formData = new EventEmitter<any>();
  surveyForm: FormGroup;
  constructor(
    public router: Router,
    private surveysService: SurveysService,
    private formBuilder: FormBuilder,
  ) { this.surveyForm = this.formBuilder.group({
    title: ['', [Validators.required]],
    summary: ['', [Validators.required]],
    disclaimer: ['', [Validators.required]],
    introduction: ['', [Validators.required]],
    });
  }

  ngOnInit(): void {
    this.surveyForm.reset();
  }
  onSurveySubmit(surveyData): void{
    this.formData.emit(surveyData);
  }
}
