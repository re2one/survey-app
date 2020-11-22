import {ChangeDetectorRef, Component, OnInit} from '@angular/core';
import {ActivatedRoute, Router} from '@angular/router';
import {QuestionsService} from '../../services/questions.service';
import {MuchoService} from '../../services/mucho.service';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';
import {AssetService} from '../../services/asset.service';
import {HttpResponse} from '@angular/common/http';
import {SurveyResponse} from '../../models/survey';
import {PuzzleAddDialogComponent, PuzzleDialogConfig} from '../puzzle-add-dialog/puzzle-add-dialog.component';
import {MatDialog} from '@angular/material/dialog';
import {PuzzleService} from '../../services/puzzle.service';

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
    private puzzleService: PuzzleService,
    private formBuilder: FormBuilder,
    public dialog: MatDialog,
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
      this.getImages();
    });
    for (let i = 0; i < 24; i++ ) {
      const piece = new PuzzlePiece(i.toString(10), parseInt(this.questionId, 10));
      this.puzzlepieces.set(i, piece);
    }
    this.load();
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
        this.getImages();
      }
    }, error => {
      console.log('error');
    });
  }
  getImages(): void {
    this.assetService.getFilenames(this.surveyId, this.questionId).subscribe( (response: HttpResponse<any>) => {
      this.filenames = response.body.filenames;
      console.log(this.filenames);
    });
    this.cdr.detectChanges();
  }
  openDialog(position: string): void{
    const config = new PuzzleDialogConfig(this.filenames, this.surveyId, this.questionId, position);
    const dialogRef = this.dialog.open(PuzzleAddDialogComponent, {
      data: config
    });
    dialogRef.afterClosed().subscribe(result => {
      const piece = this.puzzlepieces.get(result.position);
      piece.image = result.image;
      piece.empty = false;
      this.puzzlepieces.set(result.position, piece);
    });
  }
  clear(position: string): void {
    const piece = this.puzzlepieces.get(position);
    piece.image = null;
    piece.empty = true;
    this.puzzlepieces.set(position, piece);
  }
  toggleTap(position: string): void {
    const piece = this.puzzlepieces.get(position);
    piece.tapped = !piece.tapped;
    this.puzzlepieces.set(position, piece);
  }
  save(): void {
    const pieces = new Array<PuzzlePiece>();
    this.puzzlepieces.forEach((value, key) => {
      pieces.push(value);
    });
    this.puzzleService.update(this.surveyId, this.questionId, pieces).subscribe( (response: HttpResponse<any>) => {
      if (response.status === 200) {
        this.router.navigate(['/surveys/edit', this.surveyId]);
      }
    });
  }
  load(): void {
    this.puzzleService.getAll(this.questionId).subscribe((response: HttpResponse<any>) => {
      if (response.status === 200) {
        response.body.pieces.forEach( piece => {
          console.log(piece);
          this.puzzlepieces.set(parseInt(piece.position, 10), piece);
        });
      }
    });
  }
}

export class PuzzlePiece {
  constructor(
    position: string,
    questionid: number
  ) {
    this.empty = true;
    this.tapped = false;
    this.position = position;
    this.questionid = questionid;
  }
  public empty: boolean;
  public image: string;
  public position: string;
  public tapped: boolean;
  public questionid: number;
}