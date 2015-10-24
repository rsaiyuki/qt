#include "qopenglfunctions_4_2_compatibility.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QOpenGLFunctions>
#include <QOpenGLFunctions_4_2_Compatibility>
#include "_cgo_export.h"

class MyQOpenGLFunctions_4_2_Compatibility: public QOpenGLFunctions_4_2_Compatibility {
public:
};

QtObjectPtr QOpenGLFunctions_4_2_Compatibility_NewQOpenGLFunctions_4_2_Compatibility(){
	return new QOpenGLFunctions_4_2_Compatibility();
}

int QOpenGLFunctions_4_2_Compatibility_InitializeOpenGLFunctions(QtObjectPtr ptr){
	return static_cast<QOpenGLFunctions_4_2_Compatibility*>(ptr)->initializeOpenGLFunctions();
}

void QOpenGLFunctions_4_2_Compatibility_DestroyQOpenGLFunctions_4_2_Compatibility(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_4_2_Compatibility*>(ptr)->~QOpenGLFunctions_4_2_Compatibility();
}

void QOpenGLFunctions_4_2_Compatibility_GlEnd(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_4_2_Compatibility*>(ptr)->glEnd();
}

void QOpenGLFunctions_4_2_Compatibility_GlEndConditionalRender(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_4_2_Compatibility*>(ptr)->glEndConditionalRender();
}

void QOpenGLFunctions_4_2_Compatibility_GlEndList(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_4_2_Compatibility*>(ptr)->glEndList();
}

void QOpenGLFunctions_4_2_Compatibility_GlEndTransformFeedback(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_4_2_Compatibility*>(ptr)->glEndTransformFeedback();
}

void QOpenGLFunctions_4_2_Compatibility_GlFinish(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_4_2_Compatibility*>(ptr)->glFinish();
}

void QOpenGLFunctions_4_2_Compatibility_GlFlush(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_4_2_Compatibility*>(ptr)->glFlush();
}

void QOpenGLFunctions_4_2_Compatibility_GlInitNames(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_4_2_Compatibility*>(ptr)->glInitNames();
}

void QOpenGLFunctions_4_2_Compatibility_GlLoadIdentity(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_4_2_Compatibility*>(ptr)->glLoadIdentity();
}

void QOpenGLFunctions_4_2_Compatibility_GlPauseTransformFeedback(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_4_2_Compatibility*>(ptr)->glPauseTransformFeedback();
}

void QOpenGLFunctions_4_2_Compatibility_GlPopAttrib(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_4_2_Compatibility*>(ptr)->glPopAttrib();
}

void QOpenGLFunctions_4_2_Compatibility_GlPopClientAttrib(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_4_2_Compatibility*>(ptr)->glPopClientAttrib();
}

void QOpenGLFunctions_4_2_Compatibility_GlPopMatrix(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_4_2_Compatibility*>(ptr)->glPopMatrix();
}

void QOpenGLFunctions_4_2_Compatibility_GlPopName(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_4_2_Compatibility*>(ptr)->glPopName();
}

void QOpenGLFunctions_4_2_Compatibility_GlPushMatrix(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_4_2_Compatibility*>(ptr)->glPushMatrix();
}

void QOpenGLFunctions_4_2_Compatibility_GlReleaseShaderCompiler(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_4_2_Compatibility*>(ptr)->glReleaseShaderCompiler();
}

void QOpenGLFunctions_4_2_Compatibility_GlResumeTransformFeedback(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_4_2_Compatibility*>(ptr)->glResumeTransformFeedback();
}
