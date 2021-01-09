import {Component, EventEmitter, OnInit, Output} from '@angular/core';

@Component({
  selector: 'app-submit-dialog',
  templateUrl: './submit-dialog.component.html',
  styleUrls: ['./submit-dialog.component.css']
})
export class SubmitDialogComponent implements OnInit {

  @Output() shouldProceed = new EventEmitter<boolean>();

  constructor() {
  }

  ngOnInit(): void {
  }

  emitProceeding(action: boolean): void {
    this.shouldProceed.emit(action);
  }

}
