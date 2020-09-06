import {ChangeDetectorRef, Component, OnInit} from '@angular/core';
import {HttpResponse} from '@angular/common/http';
import {Survey, SurveyResponse} from '../../models/survey';
import {Router} from '@angular/router';
import {SurveysService} from '../../services/surveys.service';
import {ActivatedRoute} from '@angular/router';
import {QuestionsService} from '../../services/questions.service';

@Component({
  selector: 'app-survey-edit',
  templateUrl: './survey-edit.component.html',
  styleUrls: ['./survey-edit.component.css']
})
export class SurveyEditComponent implements OnInit {
  surveyId: string;
  survey: Survey;
  localQuestions: Map<any, any>;
  constructor(
    public router: Router,
    private surveysService: SurveysService,
    private questionsService: QuestionsService,
    private activatedRoute: ActivatedRoute,
    private cdr: ChangeDetectorRef
  ) {
    this.localQuestions = new Map();
  }

  ngOnInit(): void {
    this.activatedRoute.paramMap.subscribe(params => {
      this.surveyId = params.get('surveyId');
    });
    setTimeout(() => {
      this.questionsService.getQuestions(this.surveyId).subscribe( (response: HttpResponse<any>) => {
        if (response.status === 200) {
          response.body.questions.forEach(question => {
            this.localQuestions.set(question.ID, question);
          });
          console.log(this.localQuestions);
        }
      });
    }, 0);
  }
  onSurveySubmit(surveyData): void{
    this.surveysService.putSurvey(
      parseInt(this.surveyId, 10),
      surveyData.title,
      surveyData.summary,
      surveyData.introduction,
      surveyData.disclaimer
    ).subscribe((response: HttpResponse<SurveyResponse>) => {
      if (response.status === 200) {
        this.router.navigate(['/surveys']);
      }
    });
  }
}
