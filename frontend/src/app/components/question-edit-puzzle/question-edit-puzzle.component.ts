import {ChangeDetectorRef, Component, OnInit} from '@angular/core';
import {ActivatedRoute, Router} from '@angular/router';
import {QuestionsService} from '../../services/questions.service';
import {MuchoService} from '../../services/mucho.service';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';
import {AssetService} from '../../services/asset.service';
import {HttpResponse} from '@angular/common/http';
import {SurveyResponse} from '../../models/survey';

@Component({
  selector: 'app-question-edit-puzzle',
  templateUrl: './question-edit-puzzle.component.html',
  styleUrls: ['./question-edit-puzzle.component.css']
})
export class QuestionEditPuzzleComponent implements OnInit {

  puzzlepieces: Map<any, any>;
  questionId: string;
  surveyId: string;
  uploadForm: FormGroup;
  fileToUpload: File = null;
  filenames: Array<string>;
  constructor(
    public router: Router,
    private questionsService: QuestionsService,
    private activatedRoute: ActivatedRoute,
    private cdr: ChangeDetectorRef,
    private answersService: MuchoService,
    private assetService: AssetService,
    private formBuilder: FormBuilder,
  ) {
    this.puzzlepieces = new Map();
    this.uploadForm = this.formBuilder.group({
      file: ['', [Validators.required]],
    });
  }

  ngOnInit(): void {
    this.filenames = new Array<string>();
    this.activatedRoute.paramMap.subscribe(params => {
      this.questionId = params.get('questionId');
      this.surveyId = params.get('surveyId');
      this.assetService.getFilenames(this.surveyId, this.questionId).subscribe( (response: HttpResponse<any>) => {
        this.filenames = response.body.filenames;
        console.log(this.filenames);
      });
    });
    for (let i = 0; i < 20; i++ ) {
      this.puzzlepieces.set(i, i);
    }
    this.cdr.detectChanges();
  }
  permissionCheck(): boolean {
    const role = localStorage.getItem('role');
    return role === 'admin';
  }
  handleFileInput(files: FileList): void {
    this.fileToUpload = files.item(0);
    console.log(this.fileToUpload);
  }
  uploadFileToActivity(): void {
    this.assetService.postFile(this.fileToUpload, this.surveyId, this.questionId).subscribe((response: HttpResponse<any>) => {
      if (response.status === 200) {
        console.log('success!');
      }
    }, error => {
      console.log('error');
    });
  }
}
