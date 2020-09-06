export class Survey {
  public ID: number;
  public CreatedAt: string;
  public UpdatedAt: string;
  public DeletedAt: string;
  public title: string;
  public summary: string;
  public disclaimer: string;
  public introduction: string;
}

export class Surveys {
  public surveys: Array<Survey>;
}

export class SurveyResponse {
  public survey: Survey;
}
