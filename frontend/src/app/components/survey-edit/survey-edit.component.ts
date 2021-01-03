import {ChangeDetectorRef, Component, OnInit} from '@angular/core';
import {HttpResponse} from '@angular/common/http';
import {Survey, SurveyResponse} from '../../models/survey';
import {ActivatedRoute, Router} from '@angular/router';
import {SurveysService} from '../../services/surveys.service';
import {QuestionsService} from '../../services/questions.service';
import {BracketService} from '../../services/bracket.service';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';
import {AssetService} from '../../services/asset.service';

@Component({
  selector: 'app-survey-edit',
  templateUrl: './survey-edit.component.html',
  styleUrls: ['./survey-edit.component.css']
})
export class SurveyEditComponent implements OnInit {
  surveyId: string;
  survey: Survey;
  localQuestions: Map<any, any>;
  brackets: Array<string>;
  bracketForm: FormGroup;
  fileToUpload: File = null;
  fileSelected: boolean;

  constructor(
    public router: Router,
    private surveysService: SurveysService,
    private questionsService: QuestionsService,
    private activatedRoute: ActivatedRoute,
    private cdr: ChangeDetectorRef,
    private formBuilder: FormBuilder,
    private bracketService: BracketService,
    private assetService: AssetService,
  ) {
    this.localQuestions = new Map();
    this.bracketForm = this.formBuilder.group({
      name: ['', [Validators.required]],
    });
    this.fileSelected = false;
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
      this.getBrackets(this.surveyId);
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
  permissionCheck(): boolean {
    const role = localStorage.getItem('role');
    return role === 'admin';
  }
  delete(questionId: number): void {
    console.log(`question id to be deleted: ${questionId}`);
    this.questionsService.deleteQuestions(questionId).subscribe((response: HttpResponse<any>) => {
      if (response.status === 200) {
        this.localQuestions.delete(questionId);
        this.cdr.detectChanges();
      }
    });
  }
  moveToEditForm(questionId: number, questionType: string): void {
    console.log(questionType);
    this.router.navigate([`/questions/edit/${questionType}`, questionId, this.surveyId]);
  }
  moveToAddForm(): void {
    this.router.navigate(['/questions/add', this.surveyId]);
  }
  isFirst(first: string): boolean {
    let result = false;
    if (first === 'true') {
      result = true;
    }
    return result;
  }
  getBrackets(surveyId: string): void {
    this.bracketService.getBrackets(surveyId).subscribe((response: HttpResponse<any>) => {
      if (response.status === 200) {
        console.log(response.body.brackets);
        this.brackets = response.body.brackets;
      }
    });
  }

  onBracketSubmit(bracketData): void {
    this.bracketService.postBracket(this.surveyId, bracketData.name).subscribe((response: HttpResponse<any>) => {
      if (response.status === 200) {
        this.bracketForm.reset();
        this.getBrackets(this.surveyId);
        this.cdr.detectChanges();
      }
    });
  }

  handleFileInput(files: FileList): void {
    this.fileToUpload = files.item(0);
    this.fileSelected = true;
  }

  uploadFileToActivity(): void {
    this.assetService.postIntroduction(this.fileToUpload, this.surveyId).subscribe((response: HttpResponse<any>) => {
      if (response.status === 200) {
        this.fileSelected = false;
        this.fileToUpload = null;
        console.log('SUCCESS!');
      }
    }, error => {
      console.log('error');
    });
  }
}
