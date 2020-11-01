import {ChangeDetectorRef, Component, OnInit} from '@angular/core';
import {ActivatedRoute, Router} from '@angular/router';
import {QuestionsService} from '../../services/questions.service';
import {MuchoService} from '../../services/mucho.service';

@Component({
  selector: 'app-question-edit-puzzle',
  templateUrl: './question-edit-puzzle.component.html',
  styleUrls: ['./question-edit-puzzle.component.css']
})
export class QuestionEditPuzzleComponent implements OnInit {

  puzzlepieces: Map<any, any>;
  questionId: string;
  surveyId: string;
  constructor(
    public router: Router,
    private questionsService: QuestionsService,
    private activatedRoute: ActivatedRoute,
    private cdr: ChangeDetectorRef,
    private answersService: MuchoService
  ) {
    this.puzzlepieces = new Map();
  }

  ngOnInit(): void {
    this.activatedRoute.paramMap.subscribe(params => {
      this.questionId = params.get('questionId');
      this.surveyId = params.get('surveyId');
    });
    for (let i = 0; i < 20; i++ ) {
      this.puzzlepieces.set(i, i);
    }
    this.cdr.detectChanges();
  }
  permissionCheck(): boolean {
    const role = localStorage.getItem('role');
    return role === 'admin';
  }

}
