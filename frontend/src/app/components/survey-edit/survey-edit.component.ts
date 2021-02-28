import {ChangeDetectorRef, Component, OnInit} from '@angular/core';
import {HttpResponse} from '@angular/common/http';
import {Survey, SurveyResponse} from '../../models/survey';
import {ActivatedRoute, Router} from '@angular/router';
import {SurveysService} from '../../services/surveys.service';
import {QuestionsService} from '../../services/questions.service';
import {BracketService} from '../../services/bracket.service';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';
import {AssetService} from '../../services/asset.service';
import {DeleteDialogComponent} from '../delete-dialog/delete-dialog.component';
import {MatDialog} from '@angular/material/dialog';

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
  fileOneSelected: boolean;
  fileTwoSelected: boolean;
  fileThreeSelected: boolean;

  constructor(
    public router: Router,
    private surveysService: SurveysService,
    private questionsService: QuestionsService,
    private activatedRoute: ActivatedRoute,
    private cdr: ChangeDetectorRef,
    private formBuilder: FormBuilder,
    private bracketService: BracketService,
    private assetService: AssetService,
    public dialog: MatDialog,
  ) {
    this.localQuestions = new Map();
    this.bracketForm = this.formBuilder.group({
      name: ['', [Validators.required]],
    });
    this.fileOneSelected = false;
    this.fileTwoSelected = false;
    this.fileThreeSelected = false;
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

  openDeleteAlert(id: number): void {
    const dialogRef = this.dialog.open(DeleteDialogComponent);
    dialogRef.componentInstance.shouldProceed.subscribe(event => {
      if (event) {
        this.delete(id);
      }
    });
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

  handleFileInput(files: FileList, file: string): void {
    this.fileToUpload = files.item(0);
    switch (file) {
      case 'one':
        this.fileOneSelected = true;
        break;
      case 'two':
        this.fileTwoSelected = true;
        break;
      case 'three':
        this.fileThreeSelected = true;
        break;
      default:
        break;
    }
  }

  uploadFileToActivity(path: string): void {
    this.assetService.postAsset(this.fileToUpload, this.surveyId, path).subscribe((response: HttpResponse<any>) => {
      if (response.status === 200) {
        switch (path) {
          case '/api/assets/static/introduction/':
            this.fileOneSelected = false;
            break;
          case '/api/assets/static/termsandconditions/':
            this.fileTwoSelected = false;
            break;
          case '/api/assets/static/impressum/':
            this.fileThreeSelected = false;
            break;
          default:
            break;
        }
        this.fileToUpload = null;
        console.log('SUCCESS!');
      }
    }, error => {
      console.log('error');
    });
  }
}
