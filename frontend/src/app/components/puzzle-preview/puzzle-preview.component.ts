import {ChangeDetectorRef, Component, EventEmitter, Input, OnInit, Output} from '@angular/core';
import {Question} from '../../models/questions';
import {HttpResponse} from '@angular/common/http';
import {AssetService} from '../../services/asset.service';
import {PuzzleService} from '../../services/puzzle.service';
import {PuzzlePiece} from '../question-edit-puzzle/question-edit-puzzle.component';
import {Observable} from 'rxjs';
import {FullQuestionsService} from '../../services/full-questions.service';

@Component({
  selector: 'app-puzzle-preview',
  templateUrl: './puzzle-preview.component.html',
  styleUrls: ['./puzzle-preview.component.css']
})
export class PuzzlePreviewComponent implements OnInit {
  @Input() counter: Observable<number>;
  @Input() question: Question;
  @Output() finished = new EventEmitter<boolean>();
  filenames: Array<string>;
  puzzlepieces: Map<any, any>;

  constructor(
    private assetService: AssetService,
    private cdr: ChangeDetectorRef,
    private puzzleService: PuzzleService,
    private fqService: FullQuestionsService,
  ) {
    this.puzzlepieces = new Map();
  }
  ngOnInit(): void {
    const email = localStorage.getItem('email');
    this.getImages();
    for (let i = 0; i < 24; i++) {
      const piece = new PuzzlePiece(i.toString(10), this.question.ID);
      this.puzzlepieces.set(i, piece);
    }
    this.load(email);
    this.cdr.detectChanges();
    const myObserver = {
      next: x => {
      },
      error: err => console.error('Observer got an error: ' + err),
      complete: () => {
        this.finished.emit(false);
      },
    };
    this.counter.subscribe(myObserver);
    /*
    this.fqService.postViewed(email, this.question).subscribe((response: HttpResponse<any>) => {
      if (response.status === 200) {
        console.log('beep');
      }
    });
     */
  }

  getImages(): void {
    this.assetService.getFilenames(
      this.question.surveyid.toString(10),
      this.question.ID.toString(10)).subscribe((response: HttpResponse<any>) => {
      this.filenames = response.body.filenames;
    });
  }

  load(email: string): void {
    this.puzzleService.getAllForQuestionaire(this.question.ID.toString(10), email).subscribe(
      (response: HttpResponse<any>) => {
        if (response.status === 200) {
          response.body.pieces.forEach(piece => {
            this.puzzlepieces.set(parseInt(piece.position, 10), piece);
          });
        }
      }, (error) => {
        this.finished.emit(false);
      });
  }
}
