import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { QuestionEditPuzzleComponent } from './question-edit-puzzle.component';

describe('QuestionEditPuzzleComponent', () => {
  let component: QuestionEditPuzzleComponent;
  let fixture: ComponentFixture<QuestionEditPuzzleComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ QuestionEditPuzzleComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(QuestionEditPuzzleComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
