package controllers

import (
	"github.com/astaxie/beego"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"strings"
	"time"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) URLMapping() {
	c.Mapping("Nodes", c.Nodes)
	c.Mapping("NodePods", c.NodePods)
	c.Mapping("ContainerTerminal", c.ContainerTerminal)
}

// @router / [get]
func (c *MainController) Get() {
	c.TplName = "index.html"

}

// @Title 获取所有节点信息
// @Description 获取所有节点信息
// @Success 200 {string}
// @Failure 404 body is empty
// @router /api/nodes [get]
func (c *MainController) Nodes() {
	resp, err := Clientset.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	c.Data["json"] = resp.Items
	c.ServeJSON()
}

// @Title 获取节点上的所有容器
// @Description 获取节点上的所有容器
// @Success 200 {string}
// @Failure 404 body is empty
// @router /api/nodes/containers [get]
func (c *MainController) NodePods() {
	nodeIp := c.GetString("node")
	podList, err := Clientset.CoreV1().Pods(v1.NamespaceAll).List(metav1.ListOptions{
		FieldSelector: "status.podIP=" + nodeIp,
	})
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("nodeIp:{%v},podList:{%v},err:{%v}\n", nodeIp, podList, err)
	var mapss []interface{}
	for _, pod := range podList.Items {
		pods := pod
		var ImageID string
		var ContainerID string
		var Created int64
		var Status string
		if pods.Status.ContainerStatuses != nil {
			ImageID = pods.Status.ContainerStatuses[0].ImageID
			ContainerID = strings.Join(strings.Split(pods.Status.ContainerStatuses[0].ContainerID, "cri-o://"), "")

			if pods.Status.ContainerStatuses[0].State.Running != nil {
				Created = pods.Status.ContainerStatuses[0].State.Running.StartedAt.Unix()
			} else {
				Created = time.Now().Unix()
			}
		} else {
			ImageID = ""
			ContainerID = ""

		}
		if pods.Status.Conditions != nil {
			if pods.Status.Conditions[1].Status == "True" {
				Status = "Ready"
			} else {
				Status = pods.Status.Conditions[1].Reason
			}
		} else {
			Status = "NotReady"
		}
		maps := map[string]interface{}{
			"Name":      pods.Name,
			"Namespace": pods.Namespace,
			"NodeName":  pods.Spec.NodeName,
			"Labels":    pods.ObjectMeta.Labels,
			"SelfLink":  pods.ObjectMeta.SelfLink,
			"Uid":       pods.ObjectMeta.UID,
			"Status":    Status,
			"IP":        pods.Status.PodIP,
			"Image":     pods.Spec.Containers[0].Image,
			"AppName":   pods.Spec.Containers[0].Name,
			"ImageID":   ImageID,
			"Id":        ContainerID,
			"Created":   Created,
			"Command":   pod.Spec.Containers[0].Command,
		}
		mapss = append(mapss, maps)
	}
	c.Data["json"] = mapss
	c.ServeJSON()
}

// @router /container/terminal [get]
func (c *MainController) ContainerTerminal() {
	c.TplName = "terminal.html"
}
