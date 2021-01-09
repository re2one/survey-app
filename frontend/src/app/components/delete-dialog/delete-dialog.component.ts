import {Component, EventEmitter, OnInit, Output} from '@angular/core';

@Component({
  selector: 'app-delete-dialog',
  templateUrl: './delete-dialog.component.html',
  styleUrls: ['./delete-dialog.component.css']
})
export class DeleteDialogComponent implements OnInit {

  @Output() shouldProceed = new EventEmitter<boolean>();

  constructor() {
  }

  ngOnInit(): void {
  }

  emitProceeding(action: boolean): void {
    this.shouldProceed.emit(action);
  }

}
