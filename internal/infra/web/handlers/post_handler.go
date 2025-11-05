package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/usecases/post_usecase"
	"github.com/joaolima7/maconaria_back-end/internal/infra/web/response"
)

type PostHandler struct {
	CreatePostUseCase     *post_usecase.CreatePostUseCase
	GetAllPostsUseCase    *post_usecase.GetAllPostsUseCase
	UpdatePostByIDUseCase *post_usecase.UpdatePostByIDUseCase
	DeletePostUseCase     *post_usecase.DeletePostUseCase
}

func NewPostHandler(
	createPostUseCase *post_usecase.CreatePostUseCase,
	getAllPostsUseCase *post_usecase.GetAllPostsUseCase,
	updatePostByIDUseCase *post_usecase.UpdatePostByIDUseCase,
	deletePostUseCase *post_usecase.DeletePostUseCase,
) *PostHandler {
	return &PostHandler{
		CreatePostUseCase:     createPostUseCase,
		GetAllPostsUseCase:    getAllPostsUseCase,
		UpdatePostByIDUseCase: updatePostByIDUseCase,
		DeletePostUseCase:     deletePostUseCase,
	}
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var input post_usecase.CreatePostInputDTO

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, apperrors.NewValidationError("request", "Corpo da requisição inválido"))
		return
	}

	output, err := h.CreatePostUseCase.Execute(input)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Created(w, "Post criado com sucesso!", output)
}

func (h *PostHandler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	output, err := h.GetAllPostsUseCase.Execute()
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, "Posts buscados com sucesso!", output)
}

func (h *PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "id")
	var input post_usecase.UpdatePostByIDInputDTO

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, apperrors.NewValidationError("request", "Corpo da requisição inválido"))
		return
	}

	input.ID = postID

	output, err := h.UpdatePostByIDUseCase.Execute(input)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, "Post atualizado com sucesso!", output)
}

func (h *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "id")

	input := post_usecase.DeletePostInputDTO{ID: postID}

	if err := h.DeletePostUseCase.Execute(input); err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, "Post deletado com sucesso!", nil)
}
