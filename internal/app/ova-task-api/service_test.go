package api_test

import (
	"context"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	api "ozonva/ova-task-api/internal/app/ova-task-api"
	"ozonva/ova-task-api/internal/mocks"
	desc "ozonva/ova-task-api/pkg/api/ova-task-api"
	taskspkg "ozonva/ova-task-api/pkg/entities/tasks"
	"sort"
	"strconv"
)

var _ = Describe("Service", func() {
	var (
		ctrl          *gomock.Controller
		mockRepo      *mocks.MockRepo
		service       desc.OvaTaskApiServer
		db            map[uint64]taskspkg.Task
		autoIncrement uint64
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(ctrl)
		service = api.NewOvaTaskApi(mockRepo)
		db = make(map[uint64]taskspkg.Task, 0)
		autoIncrement = 0

		mockRepo.EXPECT().AddTasks(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, tasks []taskspkg.Task) error {
			for _, task := range tasks {
				autoIncrement++
				task.TaskId = autoIncrement
				db[autoIncrement] = task
			}
			return nil
		}).AnyTimes()
		mockRepo.EXPECT().RemoveTask(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, taskId uint64) error {
			delete(db, taskId)
			return nil
		}).AnyTimes()
		mockRepo.EXPECT().DescribeTask(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, taskId uint64) (*taskspkg.Task, error) {
			task := db[taskId]
			return &task, nil
		}).AnyTimes()
		mockRepo.EXPECT().ListTasks(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, limit, offset uint64) ([]taskspkg.Task, error) {
			keys := make([]uint64, 0, len(db))
			for k := range db {
				keys = append(keys, k)
			}
			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
			result := make([]taskspkg.Task, 0, limit)
			for _, k := range keys[offset : offset+limit] {
				result = append(result, db[k])
			}
			return result, nil

		}).AnyTimes()
	})
	Context("when add task", func() {
		It("task added", func() {
			_, err := service.CreateTaskV1(nil, &desc.CreateTaskV1Request{
				UserId:      1,
				Description: "task description",
			})
			Expect(err).Should(BeNil())
			Expect(len(db)).Should(BeEquivalentTo(1))
			task, contains := db[1]
			Expect(contains).Should(BeTrue())
			Expect(task.TaskId).Should(BeEquivalentTo(1))
			Expect(task.Description).Should(BeEquivalentTo("task description"))
		})
	})
	Context("when remove task", func() {
		It("task removed", func() {
			service.CreateTaskV1(nil, &desc.CreateTaskV1Request{
				UserId:      1,
				Description: "task description1",
			})
			service.CreateTaskV1(nil, &desc.CreateTaskV1Request{
				UserId:      1,
				Description: "task description2",
			})
			service.RemoveTasksV1(nil, &desc.RemoveTaskV1Request{TaskId: 1})
			Expect(len(db)).Should(BeEquivalentTo(1))

			task, contains := db[2]
			Expect(contains).Should(BeTrue())
			Expect(task.TaskId).Should(BeEquivalentTo(2))
			Expect(task.Description).Should(BeEquivalentTo("task description2"))
		})
	})
	Context("when list tasks", func() {
		BeforeEach(func() {
			db = make(map[uint64]taskspkg.Task, 0)
			autoIncrement = 0
			for i := 0; i < 10; i++ {
				service.CreateTaskV1(nil, &desc.CreateTaskV1Request{
					UserId:      1,
					Description: "task description" + strconv.Itoa(i),
				})
			}
		})
		It("all tasks can be listed", func() {
			response, err := service.ListTasksV1(nil, &desc.ListTasksV1Request{
				Limit:  10,
				Offset: 0,
			})
			Expect(err).Should(BeNil())
			Expect(len(response.Tasks)).Should(BeEquivalentTo(10))
			for i, task := range response.Tasks {
				Expect(task.TaskId).Should(BeEquivalentTo(i + 1))
			}
		})
		It("first task can be listed", func() {
			response, err := service.ListTasksV1(nil, &desc.ListTasksV1Request{
				Limit:  1,
				Offset: 0,
			})
			Expect(err).Should(BeNil())
			Expect(len(response.Tasks)).Should(BeEquivalentTo(1))
			task := response.Tasks[0]
			Expect(task.TaskId).Should(BeEquivalentTo(1))
		})
		It("last task can be listed", func() {
			response, err := service.ListTasksV1(nil, &desc.ListTasksV1Request{
				Limit:  1,
				Offset: 9,
			})
			Expect(err).Should(BeNil())
			Expect(len(response.Tasks)).Should(BeEquivalentTo(1))
			task := response.Tasks[0]
			Expect(task.TaskId).Should(BeEquivalentTo(10))
		})
	})
	Context("when describe task", func() {
		BeforeEach(func() {
			db = make(map[uint64]taskspkg.Task, 0)
			autoIncrement = 0
			for i := 0; i < 10; i++ {
				service.CreateTaskV1(nil, &desc.CreateTaskV1Request{
					UserId:      1,
					Description: "task description" + strconv.Itoa(i),
				})
			}
		})
		It("first task found", func() {
			response, err := service.DescribeTaskV1(nil, &desc.DescribeTaskV1Request{TaskId: 1})
			Expect(err).Should(BeNil())
			Expect(response.Task.TaskId).Should(BeEquivalentTo(1))
		})
		It("last task found", func() {
			response, err := service.DescribeTaskV1(nil, &desc.DescribeTaskV1Request{TaskId: 10})
			Expect(err).Should(BeNil())
			Expect(response.Task.TaskId).Should(BeEquivalentTo(10))
		})
	})
})
