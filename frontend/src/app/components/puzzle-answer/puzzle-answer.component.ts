import {ChangeDetectorRef, Component, Input, OnInit} from '@angular/core';
import {Question} from '../../models/questions';
import {Observable, timer} from 'rxjs';
import {map, take} from 'rxjs/operators';
import {PuzzleAddDialogComponent, PuzzleDialogConfig} from '../puzzle-add-dialog/puzzle-add-dialog.component';
import {HttpResponse} from '@angular/common/http';
import {PuzzlePiece} from '../question-edit-puzzle/question-edit-puzzle.component';
import {AssetService} from '../../services/asset.service';
import {PuzzleService} from '../../services/puzzle.service';
import {MatDialog} from '@angular/material/dialog';
import {QuestionsService} from '../../services/questions.service';
import {Router} from '@angular/router';

@Component({
  selector: 'app-puzzle-answer',
  templateUrl: './puzzle-answer.component.html',
  styleUrls: ['./puzzle-answer.component.css']
})
export class PuzzleAnswerComponent implements OnInit {
  @Input() question: Question;
  counter$: Observable<number>;
  count = 11;
  previewActive = true;
  puzzlepieces: Map<any, any>;
  filenames: Array<string>;

  constructor(
    private assetService: AssetService,
    private cdr: ChangeDetectorRef,
    private puzzleService: PuzzleService,
    public dialog: MatDialog,
    private questionsService: QuestionsService,
    private router: Router,
  ) {
    this.puzzlepieces = new Map();
    this.counter$ = timer(0, 1000).pipe(
      take(this.count),
      map(() => --this.count)
    );
    setTimeout(() => {
      this.previewActive = false;
    }, 10 * 1000);
  }

  ngOnInit(): void {
    this.getImages();
    for (let i = 0; i < 24; i++) {
      const piece = new PuzzlePiece(i.toString(10), this.question.ID);
      this.puzzlepieces.set(i, piece);
    }
    this.cdr.detectChanges();
  }

  getImages(): void {
    this.assetService.getFilenames(
      this.question.surveyid.toString(10),
      this.question.ID.toString(10)).subscribe((response: HttpResponse<any>) => {
      this.filenames = response.body.filenames;
    });
  }

  openDialog(position: string): void {
    const config = new PuzzleDialogConfig(this.filenames, this.question.surveyid.toString(10), this.question.ID.toString(10), position);
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
  }
}
