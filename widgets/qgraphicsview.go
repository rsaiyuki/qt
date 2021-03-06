package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"log"
	"unsafe"
)

type QGraphicsView struct {
	QAbstractScrollArea
}

type QGraphicsView_ITF interface {
	QAbstractScrollArea_ITF
	QGraphicsView_PTR() *QGraphicsView
}

func PointerFromQGraphicsView(ptr QGraphicsView_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsView_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsViewFromPointer(ptr unsafe.Pointer) *QGraphicsView {
	var n = new(QGraphicsView)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QGraphicsView_") {
		n.SetObjectName("QGraphicsView_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGraphicsView) QGraphicsView_PTR() *QGraphicsView {
	return ptr
}

//QGraphicsView::CacheModeFlag
type QGraphicsView__CacheModeFlag int64

const (
	QGraphicsView__CacheNone       = QGraphicsView__CacheModeFlag(0x0)
	QGraphicsView__CacheBackground = QGraphicsView__CacheModeFlag(0x1)
)

//QGraphicsView::DragMode
type QGraphicsView__DragMode int64

const (
	QGraphicsView__NoDrag         = QGraphicsView__DragMode(0)
	QGraphicsView__ScrollHandDrag = QGraphicsView__DragMode(1)
	QGraphicsView__RubberBandDrag = QGraphicsView__DragMode(2)
)

//QGraphicsView::OptimizationFlag
type QGraphicsView__OptimizationFlag int64

const (
	QGraphicsView__DontClipPainter           = QGraphicsView__OptimizationFlag(0x1)
	QGraphicsView__DontSavePainterState      = QGraphicsView__OptimizationFlag(0x2)
	QGraphicsView__DontAdjustForAntialiasing = QGraphicsView__OptimizationFlag(0x4)
	QGraphicsView__IndirectPainting          = QGraphicsView__OptimizationFlag(0x8)
)

//QGraphicsView::ViewportAnchor
type QGraphicsView__ViewportAnchor int64

const (
	QGraphicsView__NoAnchor         = QGraphicsView__ViewportAnchor(0)
	QGraphicsView__AnchorViewCenter = QGraphicsView__ViewportAnchor(1)
	QGraphicsView__AnchorUnderMouse = QGraphicsView__ViewportAnchor(2)
)

//QGraphicsView::ViewportUpdateMode
type QGraphicsView__ViewportUpdateMode int64

const (
	QGraphicsView__FullViewportUpdate         = QGraphicsView__ViewportUpdateMode(0)
	QGraphicsView__MinimalViewportUpdate      = QGraphicsView__ViewportUpdateMode(1)
	QGraphicsView__SmartViewportUpdate        = QGraphicsView__ViewportUpdateMode(2)
	QGraphicsView__NoViewportUpdate           = QGraphicsView__ViewportUpdateMode(3)
	QGraphicsView__BoundingRectViewportUpdate = QGraphicsView__ViewportUpdateMode(4)
)

func (ptr *QGraphicsView) Alignment() core.Qt__AlignmentFlag {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::alignment")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QGraphicsView_Alignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsView) BackgroundBrush() *gui.QBrush {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::backgroundBrush")
		}
	}()

	if ptr.Pointer() != nil {
		return gui.NewQBrushFromPointer(C.QGraphicsView_BackgroundBrush(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsView) CacheMode() QGraphicsView__CacheModeFlag {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::cacheMode")
		}
	}()

	if ptr.Pointer() != nil {
		return QGraphicsView__CacheModeFlag(C.QGraphicsView_CacheMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsView) DragMode() QGraphicsView__DragMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::dragMode")
		}
	}()

	if ptr.Pointer() != nil {
		return QGraphicsView__DragMode(C.QGraphicsView_DragMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsView) ForegroundBrush() *gui.QBrush {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::foregroundBrush")
		}
	}()

	if ptr.Pointer() != nil {
		return gui.NewQBrushFromPointer(C.QGraphicsView_ForegroundBrush(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsView) IsInteractive() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::isInteractive")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QGraphicsView_IsInteractive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGraphicsView) OptimizationFlags() QGraphicsView__OptimizationFlag {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::optimizationFlags")
		}
	}()

	if ptr.Pointer() != nil {
		return QGraphicsView__OptimizationFlag(C.QGraphicsView_OptimizationFlags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsView) RenderHints() gui.QPainter__RenderHint {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::renderHints")
		}
	}()

	if ptr.Pointer() != nil {
		return gui.QPainter__RenderHint(C.QGraphicsView_RenderHints(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsView) ResizeAnchor() QGraphicsView__ViewportAnchor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::resizeAnchor")
		}
	}()

	if ptr.Pointer() != nil {
		return QGraphicsView__ViewportAnchor(C.QGraphicsView_ResizeAnchor(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsView) RubberBandSelectionMode() core.Qt__ItemSelectionMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::rubberBandSelectionMode")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__ItemSelectionMode(C.QGraphicsView_RubberBandSelectionMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsView) SetAlignment(alignment core.Qt__AlignmentFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::setAlignment")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsView_SetAlignment(ptr.Pointer(), C.int(alignment))
	}
}

func (ptr *QGraphicsView) SetBackgroundBrush(brush gui.QBrush_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::setBackgroundBrush")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsView_SetBackgroundBrush(ptr.Pointer(), gui.PointerFromQBrush(brush))
	}
}

func (ptr *QGraphicsView) SetCacheMode(mode QGraphicsView__CacheModeFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::setCacheMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsView_SetCacheMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QGraphicsView) SetDragMode(mode QGraphicsView__DragMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::setDragMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsView_SetDragMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QGraphicsView) SetForegroundBrush(brush gui.QBrush_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::setForegroundBrush")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsView_SetForegroundBrush(ptr.Pointer(), gui.PointerFromQBrush(brush))
	}
}

func (ptr *QGraphicsView) SetInteractive(allowed bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::setInteractive")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsView_SetInteractive(ptr.Pointer(), C.int(qt.GoBoolToInt(allowed)))
	}
}

func (ptr *QGraphicsView) SetOptimizationFlags(flags QGraphicsView__OptimizationFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::setOptimizationFlags")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsView_SetOptimizationFlags(ptr.Pointer(), C.int(flags))
	}
}

func (ptr *QGraphicsView) SetRenderHints(hints gui.QPainter__RenderHint) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::setRenderHints")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsView_SetRenderHints(ptr.Pointer(), C.int(hints))
	}
}

func (ptr *QGraphicsView) SetResizeAnchor(anchor QGraphicsView__ViewportAnchor) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::setResizeAnchor")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsView_SetResizeAnchor(ptr.Pointer(), C.int(anchor))
	}
}

func (ptr *QGraphicsView) SetRubberBandSelectionMode(mode core.Qt__ItemSelectionMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::setRubberBandSelectionMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsView_SetRubberBandSelectionMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QGraphicsView) SetSceneRect(rect core.QRectF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::setSceneRect")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsView_SetSceneRect(ptr.Pointer(), core.PointerFromQRectF(rect))
	}
}

func (ptr *QGraphicsView) SetTransformationAnchor(anchor QGraphicsView__ViewportAnchor) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::setTransformationAnchor")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsView_SetTransformationAnchor(ptr.Pointer(), C.int(anchor))
	}
}

func (ptr *QGraphicsView) SetViewportUpdateMode(mode QGraphicsView__ViewportUpdateMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::setViewportUpdateMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsView_SetViewportUpdateMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QGraphicsView) TransformationAnchor() QGraphicsView__ViewportAnchor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::transformationAnchor")
		}
	}()

	if ptr.Pointer() != nil {
		return QGraphicsView__ViewportAnchor(C.QGraphicsView_TransformationAnchor(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsView) ViewportUpdateMode() QGraphicsView__ViewportUpdateMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::viewportUpdateMode")
		}
	}()

	if ptr.Pointer() != nil {
		return QGraphicsView__ViewportUpdateMode(C.QGraphicsView_ViewportUpdateMode(ptr.Pointer()))
	}
	return 0
}

func NewQGraphicsView2(scene QGraphicsScene_ITF, parent QWidget_ITF) *QGraphicsView {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::QGraphicsView")
		}
	}()

	return NewQGraphicsViewFromPointer(C.QGraphicsView_NewQGraphicsView2(PointerFromQGraphicsScene(scene), PointerFromQWidget(parent)))
}

func NewQGraphicsView(parent QWidget_ITF) *QGraphicsView {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::QGraphicsView")
		}
	}()

	return NewQGraphicsViewFromPointer(C.QGraphicsView_NewQGraphicsView(PointerFromQWidget(parent)))
}

func (ptr *QGraphicsView) CenterOn3(item QGraphicsItem_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::centerOn")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsView_CenterOn3(ptr.Pointer(), PointerFromQGraphicsItem(item))
	}
}

func (ptr *QGraphicsView) CenterOn(pos core.QPointF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::centerOn")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsView_CenterOn(ptr.Pointer(), core.PointerFromQPointF(pos))
	}
}

func (ptr *QGraphicsView) CenterOn2(x float64, y float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::centerOn")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsView_CenterOn2(ptr.Pointer(), C.double(x), C.double(y))
	}
}

func (ptr *QGraphicsView) EnsureVisible3(item QGraphicsItem_ITF, xmargin int, ymargin int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::ensureVisible")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsView_EnsureVisible3(ptr.Pointer(), PointerFromQGraphicsItem(item), C.int(xmargin), C.int(ymargin))
	}
}

func (ptr *QGraphicsView) EnsureVisible(rect core.QRectF_ITF, xmargin int, ymargin int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::ensureVisible")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsView_EnsureVisible(ptr.Pointer(), core.PointerFromQRectF(rect), C.int(xmargin), C.int(ymargin))
	}
}

func (ptr *QGraphicsView) EnsureVisible2(x float64, y float64, w float64, h float64, xmargin int, ymargin int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::ensureVisible")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsView_EnsureVisible2(ptr.Pointer(), C.double(x), C.double(y), C.double(w), C.double(h), C.int(xmargin), C.int(ymargin))
	}
}

func (ptr *QGraphicsView) FitInView3(item QGraphicsItem_ITF, aspectRatioMode core.Qt__AspectRatioMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::fitInView")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsView_FitInView3(ptr.Pointer(), PointerFromQGraphicsItem(item), C.int(aspectRatioMode))
	}
}

func (ptr *QGraphicsView) FitInView(rect core.QRectF_ITF, aspectRatioMode core.Qt__AspectRatioMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::fitInView")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsView_FitInView(ptr.Pointer(), core.PointerFromQRectF(rect), C.int(aspectRatioMode))
	}
}

func (ptr *QGraphicsView) FitInView2(x float64, y float64, w float64, h float64, aspectRatioMode core.Qt__AspectRatioMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::fitInView")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsView_FitInView2(ptr.Pointer(), C.double(x), C.double(y), C.double(w), C.double(h), C.int(aspectRatioMode))
	}
}

func (ptr *QGraphicsView) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::inputMethodQuery")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QGraphicsView_InputMethodQuery(ptr.Pointer(), C.int(query)))
	}
	return nil
}

func (ptr *QGraphicsView) InvalidateScene(rect core.QRectF_ITF, layers QGraphicsScene__SceneLayer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::invalidateScene")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsView_InvalidateScene(ptr.Pointer(), core.PointerFromQRectF(rect), C.int(layers))
	}
}

func (ptr *QGraphicsView) IsTransformed() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::isTransformed")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QGraphicsView_IsTransformed(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGraphicsView) ItemAt(pos core.QPoint_ITF) *QGraphicsItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::itemAt")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQGraphicsItemFromPointer(C.QGraphicsView_ItemAt(ptr.Pointer(), core.PointerFromQPoint(pos)))
	}
	return nil
}

func (ptr *QGraphicsView) ItemAt2(x int, y int) *QGraphicsItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::itemAt")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQGraphicsItemFromPointer(C.QGraphicsView_ItemAt2(ptr.Pointer(), C.int(x), C.int(y)))
	}
	return nil
}

func (ptr *QGraphicsView) Render(painter gui.QPainter_ITF, target core.QRectF_ITF, source core.QRect_ITF, aspectRatioMode core.Qt__AspectRatioMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::render")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsView_Render(ptr.Pointer(), gui.PointerFromQPainter(painter), core.PointerFromQRectF(target), core.PointerFromQRect(source), C.int(aspectRatioMode))
	}
}

func (ptr *QGraphicsView) ResetCachedContent() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::resetCachedContent")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsView_ResetCachedContent(ptr.Pointer())
	}
}

func (ptr *QGraphicsView) ResetMatrix() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::resetMatrix")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsView_ResetMatrix(ptr.Pointer())
	}
}

func (ptr *QGraphicsView) ResetTransform() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::resetTransform")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsView_ResetTransform(ptr.Pointer())
	}
}

func (ptr *QGraphicsView) Rotate(angle float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::rotate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsView_Rotate(ptr.Pointer(), C.double(angle))
	}
}

func (ptr *QGraphicsView) Scale(sx float64, sy float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::scale")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsView_Scale(ptr.Pointer(), C.double(sx), C.double(sy))
	}
}

func (ptr *QGraphicsView) Scene() *QGraphicsScene {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::scene")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQGraphicsSceneFromPointer(C.QGraphicsView_Scene(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsView) SetOptimizationFlag(flag QGraphicsView__OptimizationFlag, enabled bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::setOptimizationFlag")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsView_SetOptimizationFlag(ptr.Pointer(), C.int(flag), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QGraphicsView) SetRenderHint(hint gui.QPainter__RenderHint, enabled bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::setRenderHint")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsView_SetRenderHint(ptr.Pointer(), C.int(hint), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QGraphicsView) SetScene(scene QGraphicsScene_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::setScene")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsView_SetScene(ptr.Pointer(), PointerFromQGraphicsScene(scene))
	}
}

func (ptr *QGraphicsView) SetSceneRect2(x float64, y float64, w float64, h float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::setSceneRect")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsView_SetSceneRect2(ptr.Pointer(), C.double(x), C.double(y), C.double(w), C.double(h))
	}
}

func (ptr *QGraphicsView) SetTransform(matrix gui.QTransform_ITF, combine bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::setTransform")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsView_SetTransform(ptr.Pointer(), gui.PointerFromQTransform(matrix), C.int(qt.GoBoolToInt(combine)))
	}
}

func (ptr *QGraphicsView) Shear(sh float64, sv float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::shear")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsView_Shear(ptr.Pointer(), C.double(sh), C.double(sv))
	}
}

func (ptr *QGraphicsView) Translate(dx float64, dy float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::translate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsView_Translate(ptr.Pointer(), C.double(dx), C.double(dy))
	}
}

func (ptr *QGraphicsView) UpdateSceneRect(rect core.QRectF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::updateSceneRect")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsView_UpdateSceneRect(ptr.Pointer(), core.PointerFromQRectF(rect))
	}
}

func (ptr *QGraphicsView) DestroyQGraphicsView() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsView::~QGraphicsView")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsView_DestroyQGraphicsView(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
