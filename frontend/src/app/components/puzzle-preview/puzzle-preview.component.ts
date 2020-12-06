import {ChangeDetectorRef, Component, Input, OnInit} from '@angular/core';
import {Question} from '../../models/questions';
import {HttpResponse} from '@angular/common/http';
import {AssetService} from '../../services/asset.service';
import {PuzzleService} from '../../services/puzzle.service';
import {PuzzlePiece} from '../question-edit-puzzle/question-edit-puzzle.component';

@Component({
  selector: 'app-puzzle-preview',
  templateUrl: './puzzle-preview.component.html',
  styleUrls: ['./puzzle-preview.component.css']
})
export class PuzzlePreviewComponent implements OnInit {
  @Input() question: Question;
  filenames: Array<string>;
  puzzlepieces: Map<any, any>;

  constructor(
    private assetService: AssetService,
    private cdr: ChangeDetectorRef,
    private puzzleService: PuzzleService,
  ) {
    this.puzzlepieces = new Map();
  }

  ngOnInit(): void {
    this.getImages();
    for (let i = 0; i < 24; i++) {
      const piece = new PuzzlePiece(i.toString(10), this.question.ID);
      this.puzzlepieces.set(i, piece);
    }
    this.load();
    this.cdr.detectChanges();
  }

  getImages(): void {
    this.assetService.getFilenames(
      this.question.surveyid.toString(10),
      this.question.ID.toString(10)).subscribe((response: HttpResponse<any>) => {
      this.filenames = response.body.filenames;
    });
  }

  load(): void {
    this.puzzleService.getAll(this.question.ID.toString(10)).subscribe((response: HttpResponse<any>) => {
      if (response.status === 200) {
        response.body.pieces.forEach(piece => {
          this.puzzlepieces.set(parseInt(piece.position, 10), piece);
        });
      }
    });
  }
}
