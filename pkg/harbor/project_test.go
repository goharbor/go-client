package harbor

import (
	"context"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/project"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
	"github.com/google/uuid"
	"testing"
)

func TestSetScannerOfProject(t *testing.T) {
	uuidValue := uuid.New().String()
	projectScanner := &models.ProjectScanner{UUID: &uuidValue}
	res, err := harborClientV2.Project.SetScannerOfProject(context.TODO(), &project.SetScannerOfProjectParams{ProjectNameOrID: "library", Payload: projectScanner})
	if err != nil {
		t.Error(err.Error())
	}
	if res.IsSuccess() {
		t.Log("TEST IS OK")
	} else {
		t.Log(res.Error())
	}
}

func TestHeadProject(t *testing.T) {
	res, err := harborClientV2.Project.HeadProject(context.TODO(), &project.HeadProjectParams{ProjectName: "library"})
	if err != nil {
		t.Error(err.Error())
	}
	if res.IsSuccess() {
		t.Log("TEST IS OK")
	} else {
		t.Log(res.Error())
	}
}

func TestCreateProject(t *testing.T) {
	projectObj := &models.ProjectReq{ProjectName: "createcalico"}
	projectNew, err := harborClientV2.Project.CreateProject(context.TODO(), &project.CreateProjectParams{Project: projectObj})
	if err != nil {
		t.Error(err.Error())
	}
	if projectNew.IsSuccess() {
		t.Log("TEST IS OK")
	}
}

func TestDeleteProject(t *testing.T) {
	res, err := harborClientV2.Project.DeleteProject(context.TODO(), &project.DeleteProjectParams{ProjectNameOrID: "createcalico"})
	if err != nil {
		t.Error(err.Error())
	}
	if res.IsSuccess() {
		t.Log("TEST IS OK")
	}
}

func TestUpdateProject(t *testing.T) {
	projectOld, err := harborClientV2.Project.GetProject(context.TODO(), &project.GetProjectParams{ProjectNameOrID: "createcalico"})
	if err != nil {
		t.Error(err)
	}
	projectOld.Payload.Metadata.Public = "true"
	projectUpdate := &models.ProjectReq{Metadata: projectOld.Payload.Metadata}
	projectNew, err := harborClientV2.Project.UpdateProject(context.TODO(), &project.UpdateProjectParams{ProjectNameOrID: projectOld.Payload.Name, Project: projectUpdate})
	if err != nil {
		t.Error(err)
	}
	if projectNew.IsSuccess() {
		t.Log("TEST IS OK")
	}
}

func TestGetScannerOfProject(t *testing.T) {
	projectSanner, err := harborClientV2.Project.GetScannerOfProject(context.TODO(), &project.GetScannerOfProjectParams{ProjectNameOrID: "library"})
	if err != nil {
		t.Error(err.Error())
	}
	if projectSanner.IsSuccess() {
		t.Log(projectSanner.String())
		t.Log("TEST IS OK")
	}
}

func TestGetLogs(t *testing.T) {
	projectLogs, err := harborClientV2.Project.GetLogs(context.TODO(), &project.GetLogsParams{ProjectName: "library"})
	if err != nil {
		t.Error(err.Error())
	}
	if projectLogs.IsSuccess() {
		t.Log(projectLogs.String())
		t.Log("TEST IS OK")
	}
}

func TestGetProjectDeletable(t *testing.T) {
	projectDeletable, err := harborClientV2.Project.GetProjectDeletable(context.TODO(), &project.GetProjectDeletableParams{ProjectNameOrID: "library"})
	if err != nil {
		t.Error(err.Error())
	}
	if projectDeletable.IsSuccess() {
		t.Log(projectDeletable.String())
		t.Log("TEST IS OK")
	}
}

func TestListScannerCandidates(t *testing.T) {
	projectListScannerCandidates, err := harborClientV2.Project.ListScannerCandidatesOfProject(context.TODO(), &project.ListScannerCandidatesOfProjectParams{ProjectNameOrID: "library"})
	if err != nil {
		t.Error(err.Error())
	}
	if projectListScannerCandidates.IsSuccess() {
		t.Log(projectListScannerCandidates.String())
		t.Log("TEST IS OK")
	}
}
func TestGetProjectSummary(t *testing.T) {
	projectSummary, err := harborClientV2.Project.GetProjectSummary(context.TODO(), &project.GetProjectSummaryParams{ProjectNameOrID: "library"})
	if err != nil {
		t.Error(err.Error())
	}
	if projectSummary.IsSuccess() {
		t.Log(projectSummary.String())
		t.Log("TEST IS OK")
	}
}

func TestProjectList(t *testing.T) {
	projectList, err := harborClientV2.Project.ListProjects(context.TODO(), &project.ListProjectsParams{})
	if err != nil {
		t.Error(err.Error())
	}
	if projectList.IsSuccess() {
		t.Log(projectList.String())
		t.Log("TEST IS OK")
	}
}

func TestGetProject(t *testing.T) {
	project, err := harborClientV2.Project.GetProject(context.TODO(), &project.GetProjectParams{ProjectNameOrID: "calico"})
	if err != nil {
		t.Error(err)
	}
	if project.IsSuccess() {
		t.Log("TEST IS OK")
		t.Log(project)
	}
}
