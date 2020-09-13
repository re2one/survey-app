import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { MultipleEditComponent } from './multiple-edit.component';

describe('MultipleEditComponent', () => {
  let component: MultipleEditComponent;
  let fixture: ComponentFixture<MultipleEditComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ MultipleEditComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MultipleEditComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
