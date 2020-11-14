import {Component, EventEmitter, OnInit} from '@angular/core';
import {MatDialog, MatDialogRef, MAT_DIALOG_DATA} from '@angular/material/dialog';
import {Inject} from '@angular/core';

@Component({
  selector: 'app-puzzle-add-dialog',
  templateUrl: './puzzle-add-dialog.component.html',
  styleUrls: ['./puzzle-add-dialog.component.css']
})
export class PuzzleAddDialogComponent implements OnInit {
  constructor(
    public dialogRef: MatDialogRef<PuzzleAddDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public data: PuzzleDialogConfig,
  ) {}

  ngOnInit(): void {
  }
  emitImage(image: string, surveyId: string, questionId: string, position: string): void{
    this.dialogRef.close({image, surveyId, questionId, position});
  }

  closeDialog(): void{
    this.dialogRef.close({event: 'Cancel'});
  }

}

export class PuzzleDialogConfig {
  constructor(
    images: Array<string>,
    surveyId: string,
    questionId: string,
    position: string,
  ) {
    this.images = images;
    this.surveyId = surveyId;
    this.questionId = questionId;
    this.position = position;
  }
  public images: Array<string>;
  public surveyId: string;
  public questionId: string;
  public position: string;
}
