import {ChangeDetectorRef, Component, EventEmitter, OnInit, Output} from '@angular/core';
import {SurveysService} from '../../services/surveys.service';
import {LoginService} from '../../services/login.service';
import {Router} from '@angular/router';
import {HttpResponse} from '@angular/common/http';
import {ResultService} from '../../services/result.service';
import {MatDialog} from '@angular/material/dialog';

@Component({
  selector: 'app-surveys',
  templateUrl: './surveys.component.html',
  styleUrls: ['./surveys.component.css']
})
export class SurveysComponent implements OnInit{
  localSurveys: Map<any, any>;

  constructor(
    public dialog: MatDialog,
    private surveysService: SurveysService,
    private loginService: LoginService,
    private resultService: ResultService,
    private cdr: ChangeDetectorRef,
    public router: Router,
  ) {
    this.localSurveys = new Map();
  }

  ngOnInit(): void {
    setTimeout(() => {
      this.surveysService.getSurveys().subscribe( obj => {
        obj.surveys.forEach(survey => {
          this.localSurveys.set(survey.ID, survey);
        });
      });
    }, 0);
  }

  permissionCheck(): boolean {
    const role = localStorage.getItem('role');
    return role === 'admin';
  }
  moveToAddForm(): void {
    this.router.navigate(['/surveys/add']);
  }
  delete(id: number): void {
    this.surveysService.deleteSurvey(id).subscribe((response: HttpResponse<any>) => {
      if (response.status === 200) {
        this.localSurveys.delete(id);
        this.cdr.detectChanges();
      }
    });
  }
  moveToEditForm(surveyId: number): void {
    this.router.navigate(['/surveys/edit', surveyId]);
  }
  moveToDetails(surveyId: number): void{
    this.router.navigate(['/surveys/details', surveyId]);
  }

  getResult(surveyId: number): void {
    this.resultService.getResult(surveyId).subscribe((response: HttpResponse<any>) => {
      if (response.status === 200) {
        // console.log(response.body.result);
        const data = new Blob([response.body.result], {type: 'text/csv'});
        const url = window.URL.createObjectURL(data);
        window.open(url);
      }
    });
  }

  openDeleteAlert(id: number): void {
    const dialogRef = this.dialog.open(DeleteDialogComponent);
    dialogRef.componentInstance.shouldProceed.subscribe(event => {
      if (event) {
        this.delete(id);
      }
    });
  }
}

@Component({
  selector: 'app-delete-dialog',
  template: `
    <h1 mat-dialog-title>Attention</h1>
    <div mat-dialog-content>
      Do you really want to delete this survey?
    </div>
    <div mat-dialog-actions>
      <button mat-button mat-dialog-close (click)="emitProceeding(true)">Delete</button>
      <button mat-button mat-dialog-close (click)="emitProceeding(false)">Cancel</button>
    </div>
  `,
})
export class DeleteDialogComponent {
  @Output() shouldProceed = new EventEmitter<boolean>();

  emitProceeding(action: boolean): void {
    this.shouldProceed.emit(action);
  }
}
